// Copyright (c) Huawei Technologies Co., Ltd. 2020-2020. All rights reserved.
// Description:
// Author: l30002214
// Create: 2020/8/19

// Package base 
package base

import "sync"

const (
	HeaderDefaultLength = 4
	BodyDefaultLength   = 1024
)

// AWResult aw对象池
var aWResultPool = sync.Pool{New: func() interface{} {
	return &AWResult{
		Id:            "",
		Name:          "",
		Result:        false,
		ReqBegin:      0,
		RespTime:      0,
		TransactionId: "",
		LogInfo:       make([]interface{}, 0, HeaderDefaultLength),
		AW: AW{
			Req: Request{
				Method:    "",
				Url:       "",
				Header:    make(map[string]string, HeaderDefaultLength),
				HeaderLen: 0,
				Body:      "",
				TimeOut: TimeOut{
					Connect: 0,
					Write:   0,
					Read:    0,
				},
				End: End{
					Type:   0,
					Length: 0,
					Char:   "",
				},
			},
			Res: Response{
				Header:     make(map[string]string, HeaderDefaultLength),
				HeaderLen:  0,
				Body:       "",
				ErrReason:  "",
				StatusCode: 0,
			},
		},
	}
}}

func NewAWResult() *AWResult {
	v := aWResultPool.Get()
	if v != nil {
		return v.(*AWResult)
	}
	return &AWResult{
		Id:            "",
		Name:          "",
		Result:        false,
		ReqBegin:      0,
		RespTime:      0,
		TransactionId: "",
		LogInfo:       make([]interface{}, 0, HeaderDefaultLength),
		AW: AW{
			Req: Request{
				Method:    "",
				Url:       "",
				Header:    make(map[string]string, HeaderDefaultLength),
				HeaderLen: 0,
				Body:      "",
				TimeOut: TimeOut{
					Connect: 0,
					Write:   0,
					Read:    0,
				},
				End: End{
					Type:   0,
					Length: 0,
					Char:   "",
				},
			},
			Res: Response{
				Header:     make(map[string]string, HeaderDefaultLength),
				HeaderLen:  0,
				Body:       "",
				ErrReason:  "",
				StatusCode: 0,
			},
		},
	}
}

func (aw *AWResult) Release() {
	aw.Id = ""
	aw.Name = ""
	aw.Result = false
	aw.ReqBegin = 0
	aw.RespTime = 0
	aw.TransactionId = ""
	aw.LogInfo = aw.LogInfo[:0]

	aw.AW.Req.Method = ""
	aw.AW.Req.Url = ""
	aw.AW.Req.Header = make(map[string]string, HeaderDefaultLength)
	aw.AW.Req.HeaderLen = 0
	aw.AW.Req.Body = ""
	aw.AW.Req.TimeOut.Read = 0
	aw.AW.Req.TimeOut.Write = 0
	aw.AW.Req.TimeOut.Connect = 0
	aw.AW.Req.End.Type = 0
	aw.AW.Req.End.Length = 0
	aw.AW.Req.End.Char = ""

	aw.AW.Res.Header = make(map[string]string, HeaderDefaultLength)
	aw.AW.Res.HeaderLen = 0
	aw.AW.Res.Body = ""
	aw.AW.Res.ErrReason = ""
	aw.AW.Res.StatusCode = 0

	aWResultPool.Put(aw)
}

// 请求响应
type AWResult struct {
	Id            string
	Name          string
	Result        bool
	ReqBegin      int64
	RespTime      int64
	TransactionId string
	LogInfo       []interface{}
	AW
}

type AW struct {
	Req Request
	Res Response
}

type Request struct {
	Method    string
	Url       string
	Header    map[string]string
	HeaderLen int
	Body      string
	TimeOut
	End
}

type Response struct {
	Header     map[string]string
	HeaderLen  int
	Body       string
	ErrReason  string
	StatusCode int
}

type End struct {
	Type   uint8
	Length int
	Char   string
}

type TimeOut struct {
	Connect int64
	Write   int64
	Read    int64
}
