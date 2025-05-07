package cronpackage

import (
	"context"
	"devport/adapter/logger"
	"devport/infra/database"
	"devport/infra/file_uploader"
	"devport/usecase/user"
	"time"

	"github.com/robfig/cron/v3"
)

type RobfigCron struct {
	timeout      time.Duration
	cronText     string
	db           database.SqlInter
	logger       logger.Logger
	fileUploader file_uploader.FileUploader
}

func NewRobfigCron(timeout time.Duration, cronText string, db database.SqlInter, logger logger.Logger, fileUploader file_uploader.FileUploader) *RobfigCron {
	return &RobfigCron{
		timeout:      timeout,
		cronText:     cronText,
		db:           db,
		logger:       logger,
		fileUploader: fileUploader,
	}
}

func (r *RobfigCron) Start() {
	c := cron.New()

	// handlerをcronTextに従って追加
	if _, err := c.AddFunc(r.cronText, r.Handler); err != nil {
		r.logger.WithError(err).Errorf("error when set cron functions")
	}

	r.logger.Infof("cron start")

	c.Start()
}

func (r *RobfigCron) Handler() {
	r.logger.Infof("cron handler start")
	uc := user.NewDeleteUserImageInterator(r.db.UserRepository(), r.fileUploader, r.timeout)
	_, err := uc.Execute(context.Background(), user.DeleteUserImageInput{})
	if err != nil {
		r.logger.WithError(err).Errorf("error when delete user image")
	}
	r.logger.Infof("cron handler end")
}
