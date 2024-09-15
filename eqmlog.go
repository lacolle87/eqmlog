package eqmlog

import (
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

func LoadLogger() io.Writer {
	logFile := &lumberjack.Logger{
		Filename:   viper.GetString("logger.filename"),
		MaxSize:    viper.GetInt("logger.max_size"),
		MaxBackups: viper.GetInt("logger.max_backups"),
		MaxAge:     viper.GetInt("logger.max_age"),
		Compress:   viper.GetBool("logger.compress"),
	}

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	return multiWriter
}
