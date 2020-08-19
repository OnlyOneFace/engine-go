// Copyright (c) Huawei Technologies Co., Ltd. 2020-2020. All rights reserved.
// Description:
// Author: l30002214
// Create: 2020/8/18

// Package base 
package base

// 协议
const (
	Http = "HTTP"
	Tcp  = "TCP"
	Udp  = "UDP"
)

type BaseCase struct { // 底层结构
	PodId    int
	PodCount int
	Clients  map[string]Client
}

// 协议接口
type Client interface {
	SetLogger(Logger)
	Exec(*AWResult)
}

// 日志接口
type Logger interface {
	Debug(format string)
	Debugf(format string, v ...interface{})
	Info(format string)
	Infof(format string, v ...interface{})
	Error(format string)
	Errorf(format string, v ...interface{})
}
