package router

import (
	"context"
	"devport/adapter/api/action"
	"devport/adapter/logger"
	"devport/adapter/validator"
	"devport/domain/domain_service"
	"devport/infra/database"
	"devport/infra/email"
	"devport/infra/file_uploader"
	external_presenter "devport/presenter/external_presenter"
	user_presenter "devport/presenter/user_presenter"
	"devport/usecase/external_service_url"
	"devport/usecase/user"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type GinEngine struct {
	router       *gin.Engine
	port         Port
	ctxTimeout   time.Duration
	sql          database.SqlInter
	noSQL        database.NoSQLInter
	validator    validator.Validator
	log          logger.Logger
	email        email.Email
	fileUploader file_uploader.StorageRepositoryInter
}

func NewGinServer(
	port Port,
	t time.Duration,
	db database.SqlInter,
	validator validator.Validator,
	log logger.Logger,
	session database.NoSQLInter,
	email email.Email,
	fileUploader file_uploader.StorageRepositoryInter,
) *GinEngine {
	return &GinEngine{
		router:       gin.New(),
		port:         port,
		ctxTimeout:   t,
		sql:          db,
		noSQL:        session,
		validator:    validator,
		log:          log,
		email:        email,
		fileUploader: fileUploader,
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

	store := cookie.NewStore([]byte(os.Getenv("SESSION_SECRET")))

	router.Use(sessions.Sessions("dp_session", store))

	apiRouterGroup := router.Group("/api/v1")
	{
		apiRouterGroup.GET("/ping", e.healthCheckAction())

		apiRouterGroup.POST("/register", e.createUserAction())
		apiRouterGroup.POST("/login", e.loginUserAction())
		apiRouterGroup.GET("/is_exists", e.isExistsUserAction())

		userRouterGroup := apiRouterGroup.Group("/user/:userId")
		{
			userRouterGroup.GET("/", e.fetchUserInfoAction())

			externalServiceUrlsRouterGroup := userRouterGroup.Group("/external-service-url")
			{
				externalServiceUrlsRouterGroup.GET("/", e.findExternalServiceUrlByUser())
			}
		}

		authRouterGroup := apiRouterGroup.Group("/auth") // 認証が必要なAPI
		{
			authRouterGroup.Use(e.verifyCookieTokenAction())
			authRouterGroup.GET("/health_check", e.healthCheckAction()) // 認証状態の確認
			userAuthRouterGroup := authRouterGroup.Group("/user")       // ユーザ関連
			{
				userAuthRouterGroup.PUT("/", e.updateUserAction())
				userAuthRouterGroup.PUT("/image", e.uploadUserImageAction())
				userAuthRouterGroup.POST("/logout", e.logoutUserAction())
				userAuthRouterGroup.GET("/", e.getUserInfoAction())

				externalServiceUrlAuthRouterGroup := userAuthRouterGroup.Group("/external-service-url") // 外部サービスURL関連
				{
					externalServiceUrlAuthRouterGroup.POST("/", e.createExternalServiceUrlAction())
					externalServiceUrlAuthItemRouterGroup := externalServiceUrlAuthRouterGroup.Group("/:externalServiceUrlId")
					{
						externalServiceUrlAuthItemRouterGroup.PUT("/", e.updateExternalServiceUrlAction())
						externalServiceUrlAuthItemRouterGroup.DELETE("/", e.deleteExternalServiceUrlAction())
					}
				}

				bioAuthRouterGroup := userAuthRouterGroup.Group("/bio") // 自己紹介関連
				{
					bioAuthRouterGroup.POST("/", e.uploadBioAction())
					bioAuthRouterGroup.POST("/image", e.uploadBioImageAction())
				}
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
				user_presenter.NewCreateUserPresenter(),
				e.email,
				e.ctxTimeout,
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
				e.ctxTimeout,
			)

			act = action.NewLoginUserAction(uc, e.validator, e.log)
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
				e.ctxTimeout,
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
				e.fileUploader.BioSentenceStorageRepository(),
				user_presenter.NewGetUserInfoPresenter(),
				e.ctxTimeout,
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

func (e *GinEngine) isExistsUserAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = user.NewIsExistsUserInteractor(
				e.sql.UserRepository(),
				user_presenter.NewIsExistsUserPresenter(),
				e.ctxTimeout,
			)

			act = action.NewIsExistsUserAction(uc, e.validator, e.log)
		)

		act.Execute(c.Writer, c.Request, c)
	}
}

func (e *GinEngine) fetchUserInfoAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = user.NewGetUserInfoInterator(
				e.sql.UserRepository(),
				e.fileUploader.BioSentenceStorageRepository(),
				user_presenter.NewGetUserInfoPresenter(),
				e.ctxTimeout,
			)

			act = action.NewFetchUserInfoAction(uc, e.validator, e.log)
		)

		act.Execute(c.Writer, c.Request, c)
	}
}

func (e *GinEngine) updateUserAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = user.NewUpdateUserInterator(
				e.sql.UserRepository(),
				e.sql.ExternalServiceUrlsRepository(),
				e.fileUploader.BioSentenceStorageRepository(),
				user_presenter.NewUpdateUserPresenter(),
				e.ctxTimeout,
			)

			act = action.NewUpdateUserAction(uc, e.validator, e.log)
		)

		act.Execute(c.Writer, c.Request, c)
	}
}

func (e *GinEngine) uploadUserImageAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = user.NewUploadUserImageInterator(
				e.sql.UserRepository(),
				e.fileUploader.IconImageStorageRepository(),
				e.fileUploader.HeaderImageStorageRepository(),
				user_presenter.NewUploadUserImagePresenter(),
				e.ctxTimeout,
			)

			act = action.NewUploadUserImageAction(uc, e.validator, e.log)
		)

		act.Execute(c.Writer, c.Request, c)
	}
}

func (e *GinEngine) createExternalServiceUrlAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = external_service_url.NewCreateExternalServiceUrlInteractor(
				e.sql.ExternalServiceUrlsRepository(),
				e.sql.UserRepository(),
				external_presenter.NewCreateExternalServiceUrlPresenter(),
				e.ctxTimeout,
			)

			act = action.NewCreateExternalServiceUrlAction(uc, e.validator, e.log)
		)

		act.Execute(c.Writer, c.Request, c)
	}
}

func (e *GinEngine) updateExternalServiceUrlAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = external_service_url.NewUpdateExternalServiceUrlInteractor(
				e.sql.ExternalServiceUrlsRepository(),
				e.sql.UserRepository(),
				external_presenter.NewUpdateExternalServiceUrlPresenter(),
				e.ctxTimeout,
			)

			act = action.NewUpdateExternalServiceUrlAction(uc, e.validator, e.log)
		)

		act.Execute(c.Writer, c.Request, c)
	}
}

func (e *GinEngine) deleteExternalServiceUrlAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = external_service_url.NewDeleteExternalServiceUrlInteractor(
				e.sql.ExternalServiceUrlsRepository(),
				e.sql.UserRepository(),
				external_presenter.NewDeleteExternalServiceUrlPresenter(),
				e.ctxTimeout,
			)

			act = action.NewDeleteExternalServiceUrlAction(uc, e.validator, e.log)
		)

		act.Execute(c.Writer, c.Request, c)
	}
}

func (e *GinEngine) findExternalServiceUrlByUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = external_service_url.NewFindByUserExternalServiceUrlInterator(
				e.sql.UserRepository(),
				e.sql.ExternalServiceUrlsRepository(),
				external_presenter.NewFindByUserExternalServiceUrlPresenter(),
				e.ctxTimeout,
			)

			act = action.NewFindByUserExternalServiceUrlAction(uc, e.validator, e.log)
		)

		act.Execute(c.Writer, c.Request, c)
	}
}

func (e *GinEngine) uploadBioAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = user.NewUploadBioInteractor(
				domain_service.NewBioService(e.fileUploader.BioImageStorageRepository()),
				e.sql.UserRepository(),
				e.sql.BioImagesRepository(),
				e.fileUploader.BioSentenceStorageRepository(),
				e.fileUploader.BioImageStorageRepository(),
				user_presenter.NewUploadBioPresenter(),
				e.ctxTimeout,
			)

			act = action.NewUploadBioAction(uc, e.validator, e.log)
		)

		act.Execute(c.Writer, c.Request, c)
	}
}

func (e *GinEngine) uploadBioImageAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = user.NewUploadBioImageInteractor(
				e.sql.UserRepository(),
				e.sql.BioImagesRepository(),
				e.fileUploader.BioImageStorageRepository(),
				e.sql.BioImagesRepository(),
				user_presenter.NewUploadBioImagePresenter(),
				e.ctxTimeout,
			)

			act = action.NewUploadBioImageAction(uc, e.validator, e.log)
		)

		act.Execute(c.Writer, c.Request, c)
	}
}
