package logger

import (
	"os"
	"time"

	"github.com/SilleCao/golang/go-micro-demo/internal/pkg/common"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

var Logger zerolog.Logger

func InitLog() {
	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}
	// multi := zerolog.MultiLevelWriter(consoleWriter, os.Stdout)
	multi := zerolog.MultiLevelWriter(consoleWriter)
	Logger = zerolog.New(multi).With().Stack().Timestamp().Logger()
}

func Info(msg string, ctx *gin.Context) {
	Logger.Info().Str(common.GetReqIdKey(), common.GetReqIdValue(ctx)).Msg(msg)
}

func InfoWithErr(msg string, ctx *gin.Context, err error) {
	Logger.Info().Err(err).Str(common.GetReqIdKey(), common.GetReqIdValue(ctx)).Msg(msg)
}

func Err(msg string, ctx *gin.Context, err error) {
	if err != nil {
		Logger.Error().Err(err).Str(common.GetReqIdKey(), common.GetReqIdValue(ctx)).Msg(msg)
		return
	}
	Logger.Error().Str(common.GetReqIdKey(), common.GetReqIdValue(ctx)).Msg(msg)
}

func Warn(msg string, ctx *gin.Context, err error) {
	if err != nil {
		Logger.Warn().Err(err).Str(common.GetReqIdKey(), common.GetReqIdValue(ctx)).Msg(msg)
		return
	}
	Logger.Warn().Str(common.GetReqIdKey(), common.GetReqIdValue(ctx)).Msg(msg)
}
