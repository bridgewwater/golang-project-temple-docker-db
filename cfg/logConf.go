package cfg

import (
	"github.com/bridgewwater/golang-project-temple-docker-db/pkg/zlog"
)

// Initialization log
func (c *ConfFile) initLog() error {
	logger := zlog.LoggerByViper(c.WriteLogFile)
	zlog.NewZapLog(logger, logger.Sugar())
	return nil
}
