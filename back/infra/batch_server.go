package infra

import (
	"devport/adapter/logger"
	"devport/adapter/repository"
	cronpackage "devport/infra/cron"
	"devport/infra/database"
	"devport/infra/file_uploader"
	"devport/infra/log"
	"time"
)

type BatchServerConfig struct {
	appName      string
	ctxTimeout   time.Duration
	logger       logger.Logger
	dbSql        repository.SQL
	fileUploader file_uploader.FileUploader
	cronText     string
	cs           cronpackage.CronServer
}

func NewBatchServerConfig() *BatchServerConfig {
	return &BatchServerConfig{}
}

func (c *BatchServerConfig) AppName(appName string) *BatchServerConfig {
	c.appName = appName

	return c
}

func (c *BatchServerConfig) CtxTimeout(ctxTimeout time.Duration) *BatchServerConfig {
	c.ctxTimeout = ctxTimeout
	return c
}

func (c *BatchServerConfig) CronText(ct string) *BatchServerConfig {
	c.cronText = ct
	return c
}

func (c *BatchServerConfig) LoggingTool(instance int) *BatchServerConfig {
	l, err := log.NewLoggerFactory(instance)
	if err != nil {
		panic(err)
	}
	c.logger = l
	return c
}

func (c *BatchServerConfig) DB(instance int) *BatchServerConfig {
	db, err := database.NewDatabaseSqlFactory(instance)

	if err != nil {
		panic(err) // TODO: loggerの追加
	}

	c.dbSql = db

	return c
}

func (c *BatchServerConfig) FileUploader(instance int) *BatchServerConfig {
	uploader, err := file_uploader.NewFileUploaderFactory(instance)

	if err != nil {
		panic(err)
	}

	c.fileUploader = uploader

	return c
}

func (c *BatchServerConfig) Cron(instance int) *BatchServerConfig {
	cr, err := cronpackage.NewCronFactory(instance, c.ctxTimeout, c.cronText, c.dbSql, c.logger, c.fileUploader)

	if err != nil {
		panic(err)
	}

	c.cs = cr

	return c
}

func (c *BatchServerConfig) Start() {
	c.cs.Start()
}
