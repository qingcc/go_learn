package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/polaris1119/logger"
	"os"
)

func GetLogger(c *gin.Context) *logger.Logger {
	if c == nil {
		return logger.New(os.Stdout)
	}

	_logger, ok := c.Value("logger").(*logger.Logger)
	if ok {
		return _logger
	}

	return logger.New(os.Stdout)
}
