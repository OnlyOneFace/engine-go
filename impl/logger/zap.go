// Copyright (c) Huawei Technologies Co., Ltd. 2020-2020. All rights reserved.
// Description:
// Author: l30002214
// Create: 2020/8/19

// Package logger 
package logger

import (
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	DebugLevel  = "DEBUG"
	InfoLevel   = "INFO"
	WarnLevel   = "WARN"
	ErrorLevel  = "ERROR"
	DpanicLevel = "DPANIC"
	PanicLevel  = "PANIC"
	FatalLevel  = "FATAL"
)

type Encoder func(zapcore.EncoderConfig) zapcore.Encoder

func NewZap(level string, encoderFunc Encoder, w zapcore.WriteSyncer, fields ...zap.Field) *zap.Logger {
	core := zapcore.NewCore(
		encoderFunc(newEncoderConfig()),
		zap.CombineWriteSyncers(w),
		newLevel(level),
	).With(fields) //自带node 信息
	//大于error增加堆栈信息
	return zap.New(core).WithOptions(zap.AddCaller(), zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.DPanicLevel))
}

func newEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:    "message",
		LevelKey:      "level",
		TimeKey:       "logTime",
		NameKey:       "logger",
		CallerKey:     "caller",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		EncodeTime: func(i time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(i.Local().Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
}

func newLevel(level string) zapcore.Level {
	var l zapcore.Level
	switch strings.ToUpper(level) {
	case DebugLevel:
		l = zap.DebugLevel
	case InfoLevel:
		l = zap.InfoLevel
	case WarnLevel:
		l = zap.WarnLevel
	case ErrorLevel:
		l = zap.ErrorLevel
	case DpanicLevel:
		l = zap.DPanicLevel
	case PanicLevel:
		l = zap.PanicLevel
	case FatalLevel:
		l = zap.FatalLevel
	default:
		l = zap.DebugLevel
	}
	return l
}
