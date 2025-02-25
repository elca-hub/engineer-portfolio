package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"devport/controller"
	"devport/domain/repository"
	"devport/infra/token_auth"
	user_presenter "devport/presenter/user_presenter"
	"devport/usecase/user"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type GinEngine struct {
	router         *gin.Engine
	port           Port
	ctxTimeout     time.Duration
	sql            repository.SQL
	noSQL          repository.NoSQL
	userController controller.UserController
}

func NewGinServer(
	port Port,
	t time.Duration,
	db repository.SQL,
	session repository.NoSQL,
) *GinEngine {
	return &GinEngine{
		router:     gin.New(),
		port:       port,
		ctxTimeout: t,
		sql:        db,
		noSQL:      session,
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
	router.GET("/ping", e.healthCheckAction())

	apiRouterGroup := router.Group("/api/v1")
	{
		apiRouterGroup.POST("/signup", func(c *gin.Context) {
			res := e.userController.CreateUser(c.PostForm("name"), c.PostForm("email"), c.PostForm("password"))

			c.JSON(res.StatusCode, res)
		})
		apiRouterGroup.POST("/login", func(c *gin.Context) {
			res := e.userController.LoginUser(c.PostForm("email"), c.PostForm("password"))

			token_auth.SetToken(c.Writer, res.Data["Email"].(string))

			c.JSON(res.StatusCode, res)
		})

		userRouterGroup := apiRouterGroup.Group("/user")
		{
			userRouterGroup.GET("/", e.getUserInfoAction())
		}
		apiRouterGroup.GET("/verification/email", e.verificationEmailAction())
	}
}

func (e *GinEngine) healthCheckAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
		})
	}
}

func (e *GinEngine) verificationEmailAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		uc := user.NewVerificationEmailInterator(
			e.sql.UserRepository(),
			e.noSQL.UserRepository(),
			user_presenter.NewVerificationEmailPresenter(),
		)

		token, err := uc.Execute(user.VerificationEmailInput{
			Token: c.DefaultQuery("token", ""),
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"err":  err.Error(),
			})
			return
		}

		token_auth.SetToken(c.Writer, token.Token)

		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
		})
	}
}

func (e *GinEngine) getUserInfoAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		uc := user.NewGetUserInfoInterator(
			e.sql.UserRepository(),
			e.noSQL.UserRepository(),
			user_presenter.NewGetUserInfoPresenter(),
		)

		userOutput, err := uc.Execute(user.GetUserInfoInput{
			Token: token_auth.GetToken(c.Request),
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"err":  err.Error(),
			})
			return
		}

		token_auth.SetToken(c.Writer, userOutput.Token)

		c.JSON(http.StatusOK, gin.H{
			"code":  http.StatusOK,
			"email": userOutput.Email,
		})
	}
}

func (e *GinEngine) loginAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		uc := user.NewLoginUserInterator(
			e.sql.UserRepository(),
			e.noSQL.UserRepository(),
			user_presenter.NewLoginUserPresenter(),
		)

		userOutput, err := uc.Execute(user.LoginUserInput{
			Email:    c.PostForm("email"),
			Password: c.PostForm("password"),
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"err":  err.Error(),
			})
			return
		}

		token_auth.SetToken(c.Writer, userOutput.Token)

		c.JSON(http.StatusOK, gin.H{
			"code":           http.StatusOK,
			"user_presenter": userOutput,
		})
	}
}
