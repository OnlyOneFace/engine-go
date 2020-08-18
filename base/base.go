// Copyright (c) Huawei Technologies Co., Ltd. 2020-2020. All rights reserved.
// Description:
// Author: l30002214
// Create: 2020/8/18

// Package base 
package base

type BaseCase struct { // 底层结构
	Logger
	Client
}

// 协议接口
type Client interface {
	Exec(*AWResult)
}

// 日志接口
type Logger interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
}

// 请求响应
type AWResult struct {
	Id       string
	Name     string
	Result   bool
	ReqBegin int64
	RespTime int64
	AW
	LogInfo       []string
	TransactionId string
}

type AW struct {
	Req Request
	res Response
}

type Request struct {
	Method    string
	Url       string
	Header    map[string]string
	HeaderLen int
	Body      []byte
}

type Response struct {
	Header     map[string]string
	HeaderLen  int
	Body       []byte
	ErrReason  string
	StatusCode int
	End
}

type End struct {
	Type   uint8
	Length int
	Char   string
}
