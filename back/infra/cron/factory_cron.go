package cronpackage

import (
	"devport/adapter/logger"
	"devport/infra/database"
	"devport/infra/file_uploader"
	"fmt"
	"time"
)

type CronServer interface {
	Start()
}

const (
	InstanceCron int = iota
)

func NewCronFactory(
	instance int,
	timeout time.Duration,
	cronText string,
	db database.SqlInter,
	logger logger.Logger,
	fileUploader file_uploader.FileUploader,
) (CronServer, error) {
	switch instance {
	case InstanceCron:
		return NewRobfigCron(timeout, cronText, db, logger, fileUploader), nil
	default:
		return nil, fmt.Errorf("instance not exist")
	}
}
