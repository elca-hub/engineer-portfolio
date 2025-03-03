package infra

import (
	"devport/adapter/logger"
	"devport/adapter/validator"
	"devport/domain/repository"
	"devport/infra/database"
	"devport/infra/email"
	"devport/infra/log"
	"devport/infra/router"
	"devport/infra/validation"
	"strconv"
	"time"
)

type HttpServerConfig struct {
	appName       string
	ctxTimeout    time.Duration
	validator     validator.Validator
	logger        logger.Logger
	dbSql         repository.SQL
	dbNoSql       repository.NoSQL
	webServer     router.Server
	webServerPort router.Port
	email         email.Email
}

func NewHttpServerConfig() *HttpServerConfig {
	return &HttpServerConfig{}
}

func (c *HttpServerConfig) Name(appName string) *HttpServerConfig {
	c.appName = appName

	return c
}

func (c *HttpServerConfig) ContextTimeout(timeout time.Duration) *HttpServerConfig {
	c.ctxTimeout = timeout

	return c
}

func (c *HttpServerConfig) DbSql(instance int) *HttpServerConfig {
	db, err := database.NewDatabaseSqlFactory(instance)

	if err != nil {
		panic(err) // TODO: loggerの追加
	}

	c.dbSql = db

	return c
}

func (c *HttpServerConfig) DbNoSql(instance int) *HttpServerConfig {
	db, err := database.NewDatabaseNoSqlFactory(instance)
	if err != nil {
		panic(err) // TODO: loggerの追加
	}
	c.dbNoSql = db

	return c
}

func (c *HttpServerConfig) WebServerPort(port string) *HttpServerConfig {
	p, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		panic(err) // TODO: loggerの追加
	}

	c.webServerPort = router.Port(p)
	return c
}

func (c *HttpServerConfig) Validator(instance int) *HttpServerConfig {
	v, err := validation.NewValidationFactory(instance)

	if err != nil {
		panic(err)
	}

	c.validator = v
	return c
}

func (c *HttpServerConfig) Logger(instance int) *HttpServerConfig {
	l, err := log.NewLoggerFactory(instance)
	if err != nil {
		panic(err)
	}
	c.logger = l
	return c
}

func (c *HttpServerConfig) Email() *HttpServerConfig {
	c.email = email.NewSmtp()

	return c
}

func (c *HttpServerConfig) WebServer(instance int) *HttpServerConfig {
	s, err := router.NewWebServerFactory(
		instance,
		c.webServerPort,
		c.ctxTimeout,
		c.dbSql,
		c.dbNoSql,
		c.validator,
		c.logger,
		c.email,
	)

	if err != nil {
		panic(err) // TODO: loggerの追加
	}

	c.webServer = s

	return c
}

func (c *HttpServerConfig) Start() {
	c.webServer.Listen()
}
