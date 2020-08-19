// Copyright (c) Huawei Technologies Co., Ltd. 2020-2020. All rights reserved.
// Description:
// Author: l30002214
// Create: 2020/8/18

// Package base 
package base

// 协议
const (
	Http Protocol = "HTTP"
	Tcp  Protocol = "TCP"
	Udp  Protocol = "UDP"
)

// 已经实现的协议池
var Clients = make(map[Protocol]Client)

type Protocol string

func Register(pt Protocol, c Client) {
	Clients[pt] = c
}

//协议

type BaseCase struct { // 底层结构
	PodId    int
	PodCount int
	Clients  map[Protocol]Client
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
