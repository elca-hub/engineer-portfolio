package router

import (
	"context"
	"devport/adapter/api/action"
	"devport/adapter/logger"
	"devport/adapter/validator"
	"devport/domain/repository"
	user_presenter "devport/presenter/user_presenter"
	"devport/usecase/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type GinEngine struct {
	router     *gin.Engine
	port       Port
	ctxTimeout time.Duration
	sql        repository.SQL
	noSQL      repository.NoSQL
	validator  validator.Validator
	log        logger.Logger
}

func NewGinServer(
	port Port,
	t time.Duration,
	db repository.SQL,
	validator validator.Validator,
	log logger.Logger,
	session repository.NoSQL,
) *GinEngine {
	return &GinEngine{
		router:     gin.New(),
		port:       port,
		ctxTimeout: t,
		sql:        db,
		noSQL:      session,
		validator:  validator,
		log:        log,
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
	apiRouterGroup := router.Group("/api/v1")
	{
		apiRouterGroup.GET("/ping", e.healthCheckAction())

		apiRouterGroup.POST("/signup", e.createUserAction())
		apiRouterGroup.POST("/login", e.loginUserAction())
		apiRouterGroup.GET("/verification/email", e.verificationEmailAction())

		//userRouterGroup := apiRouterGroup.Group("/user")
		//{
		//	userRouterGroup.GET("/", e.getUserInfoAction())
		//}
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
