// Copyright (c) Huawei Technologies Co., Ltd. 2020-2020. All rights reserved.
// Description:
// Author: l30002214
// Create: 2020/8/19

// Package logger 
package logger

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewZapLogger(level string, enableDebugLog bool) *ZapLogger {
	var w io.Writer
	if enableDebugLog {
		w = os.Stdout
	} else {
		w = &lumberjack.Logger{
			Filename:   fmt.Sprintf("./logs/debug_go.%s.log", time.Now().Format("2006-01-02")),
			MaxSize:    5, //单个日志文件最大MaxSize*M大小 // megabytes
			MaxBackups: 5, // 保存历史日志文件的最大个数
			MaxAge:     2, //days
			Compress:   false,
		}
	}
	return &ZapLogger{
		NewZap(
			strings.ToLower(level),
			zapcore.NewConsoleEncoder, zapcore.AddSync(w),
		),
	}
}

type ZapLogger struct {
	*zap.Logger
}

func (z *ZapLogger) Debug(format string) {
	z.Logger.Debug(format)
}

func (z *ZapLogger) Debugf(format string, v ...interface{}) {
	z.Sugar().Debugf(format, v...)
}

func (z *ZapLogger) Info(format string) {
	z.Logger.Info(format)
}

func (z *ZapLogger) Infof(format string, v ...interface{}) {
	z.Sugar().Infof(format, v...)
}

func (z *ZapLogger) Error(format string) {
	z.Logger.Error(format)
}

func (z *ZapLogger) Errorf(format string, v ...interface{}) {
	z.Sugar().Errorf(format, v...)
}
