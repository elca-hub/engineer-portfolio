package router

import (
	"context"
	"devport/adapter/api/action"
	"devport/adapter/logger"
	"devport/adapter/validator"
	"devport/domain/repository"
	"devport/infra/email"
	user_presenter "devport/presenter/user_presenter"
	"devport/usecase/user"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type GinEngine struct {
	router     *gin.Engine
	port       Port
	ctxTimeout time.Duration
	sql        repository.SQL
	noSQL      repository.NoSQL
	validator  validator.Validator
	log        logger.Logger
	email      email.Email
}

func NewGinServer(
	port Port,
	t time.Duration,
	db repository.SQL,
	validator validator.Validator,
	log logger.Logger,
	session repository.NoSQL,
	email email.Email,
) *GinEngine {
	return &GinEngine{
		router:     gin.New(),
		port:       port,
		ctxTimeout: t,
		sql:        db,
		noSQL:      session,
		validator:  validator,
		log:        log,
		email:      email,
	}
}

func (e *GinEngine) Listen() {
	gin.SetMode(gin.ReleaseMode)
	gin.Recovery()

	e.setupRouter(e.router)

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         fmt.Sprintf(":%d", e.port),
		Handler:      e.router,
	}
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		// TODO: logの追加
		fmt.Println("web server running!")
		if err := server.ListenAndServe(); err != nil {
			// TODO: errorlog追加
			fmt.Println("web server stopped")
		}
	}()
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), e.ctxTimeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
}

func (e *GinEngine) setupRouter(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Set-Cookie"},
		AllowCredentials: true,
	}))
	apiRouterGroup := router.Group("/api/v1")
	{
		apiRouterGroup.GET("/ping", e.healthCheckAction())

		apiRouterGroup.POST("/signup", e.createUserAction())
		apiRouterGroup.POST("/login", e.loginUserAction())
		apiRouterGroup.GET("/verification/email", e.verificationEmailAction())
		apiRouterGroup.POST("/logout", e.logoutUserAction())

		authRouterGroup := apiRouterGroup.Group("/auth")
		{
			authRouterGroup.Use(e.verifyCookieTokenAction())
			userRouterGroup := authRouterGroup.Group("/user")
			{
				userRouterGroup.GET("/", e.getUserInfoAction())
			}
		}
	}
}

func (e *GinEngine) healthCheckAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
		})
	}
}

func (e *GinEngine) createUserAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = user.NewCreateUserInterator(
				e.sql.UserRepository(),
				e.noSQL.UserRepository(),
				e.email,
			)

			act = action.NewCreateUserAction(uc, e.validator, e.log)
		)

		act.Execute(c.Writer, c.Request)
	}
}

func (e *GinEngine) loginUserAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = user.NewLoginUserInterator(
				e.sql.UserRepository(),
				e.noSQL.UserRepository(),
				user_presenter.NewLoginUserPresenter(),
			)

			act = action.NewLoginUserAction(uc, e.validator, e.log)
		)

		act.Execute(c.Writer, c.Request)
	}
}

func (e *GinEngine) verificationEmailAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = user.NewVerificationEmailInterator(
				e.sql.UserRepository(),
				e.noSQL.UserRepository(),
				user_presenter.NewVerificationEmailPresenter(),
			)

			act = action.NewVerifyEmailAction(uc, e.validator, e.log)
		)

		act.Execute(c.Writer, c.Request)
	}
}

func (e *GinEngine) verifyCookieTokenAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = user.NewVerifyCookieTokenInterator(
				e.sql.UserRepository(),
				e.noSQL.UserRepository(),
				user_presenter.NewVerifyCookieTokenPresenter(),
			)

			act = action.NewVerifyCookieTokenAction(uc, e.validator, e.log)
		)

		act.Execute(c.Writer, c.Request, c)

		if c.Writer.Status() != http.StatusOK {
			c.Abort()
		}
	}
}

func (e *GinEngine) getUserInfoAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = user.NewGetUserInfoInterator(
				e.sql.UserRepository(),
				e.noSQL.UserRepository(),
				user_presenter.NewGetUserInfoPresenter(),
			)

			act = action.NewGetUserAction(uc, e.validator, e.log)
		)

		act.Execute(c.Writer, c.Request, c)
	}
}

func (e *GinEngine) logoutUserAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = user.NewLogoutUserInterator(
				e.noSQL.UserRepository(),
				user_presenter.NewLogoutUserPresenter(),
			)

			act = action.NewLogoutUserAction(uc, e.validator, e.log)
		)

		act.Execute(c.Writer, c.Request)
	}
}
