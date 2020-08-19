// Copyright (c) Huawei Technologies Co., Ltd. 2020-2020. All rights reserved.
// Description:
// Author: l30002214
// Create: 2020/8/19

// Package logger 
package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
)

func NewRus(level string, format logrus.Formatter, w io.Writer, fields logrus.Fields) *logrus.Entry {
	return (&logrus.Logger{
		Out:          w,
		Hooks:        make(logrus.LevelHooks),
		Formatter:    format,
		ReportCaller: true,
		Level:        getLevel(level),
		ExitFunc:     os.Exit,
	}).WithFields(fields)
}

func getLevel(level string) logrus.Level {
	l, err := logrus.ParseLevel(strings.ToLower(level))
	if err != nil {
		l = logrus.DebugLevel
	}
	return l
}
