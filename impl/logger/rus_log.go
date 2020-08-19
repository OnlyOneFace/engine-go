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
		&logrus.TextFormatter{},
		SetWriter(enableDebugLog),
		map[string]interface{}{})}
}

type RusLogger struct {
	*logrus.Entry
}

func (r *RusLogger) Debug(format string) {
	r.Entry.Debug(format)
}

func (r *RusLogger) Info(format string) {
	r.Entry.Info(format)
}

func (r *RusLogger) Error(format string) {
	r.Entry.Error(format)
}
