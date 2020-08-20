// Copyright (c) Huawei Technologies Co., Ltd. 2020-2020. All rights reserved.
// Description:
// Author: l30002214
// Create: 2020/8/19

// Package logger 
package logger

import "github.com/sirupsen/logrus"

func NewRusLogger(level string, enableDebugLog bool) *RusLogger {
	return &RusLogger{NewRus(
		level,
		&logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		},
		SetWriter(enableDebugLog),
		map[string]interface{}{})}
}

type RusLogger struct {
	*logrus.Logger
}

func (r *RusLogger) Debug(format string) {
	r.Logger.Debug(format)
}

func (r *RusLogger) Info(format string) {
	r.Logger.Info(format)
}

func (r *RusLogger) Error(format string) {
	r.Logger.Error(format)
}
