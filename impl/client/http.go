// Copyright (c) Huawei Technologies Co., Ltd. 2020-2020. All rights reserved.
// Description:
// Author: l30002214
// Create: 2020/8/18

// Package impl 
package client

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/valyala/fasthttp"

	"github.com/OnlyOneFace/engine-go/base"
	"github.com/OnlyOneFace/engine-go/impl/logger"
	"github.com/OnlyOneFace/engine-go/util"
)

//插件最终包含的是base

const (
	DefaultContentTypeKey   = fasthttp.HeaderContentType
	DefaultContentTypeValue = "application/json;charset=utf-8"

	maxRedirectsCount = 3
)

var (
	fhConnectTimeout      = 30 * time.Second
	fhWriteTimeOut        = 60 * time.Second
	fhReadTimeout         = 60 * time.Second
	fhMaxIdleConnDuration = 60 * time.Second
	fastHttpClient        = &fasthttp.Client{
		NoDefaultUserAgentHeader: true,
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
			MaxVersion:         tls.VersionTLS12,
			MinVersion:         tls.VersionTLS12},
		MaxConnsPerHost: 60000,
		Dial: func(addr string) (net.Conn, error) {
			return (&fasthttp.TCPDialer{Concurrency: 60000}).DialTimeout(addr, fhConnectTimeout)
		},
		WriteTimeout:                  fhWriteTimeOut,
		ReadTimeout:                   fhReadTimeout,
		MaxIdleConnDuration:           fhMaxIdleConnDuration,
		DisableHeaderNamesNormalizing: true}
)

func init() {
	base.Register(base.Http, &FastHttp{
		Logger: logger.NewZapLogger("DEBUG", true),
	})
	log.Println("register http finish!")
}

type FastHttp struct {
	base.Logger
}

func (f *FastHttp) SetLogger(logger base.Logger) {
	f.Logger = logger
}

func (f *FastHttp) Exec(aw *base.AWResult) {
	start := time.Now()
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		aw.ReqBegin = start.UnixNano() / 1e6
		aw.RespTime = time.Since(start).Nanoseconds() / 1e6
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()

	req.Header.DisableNormalizing()
	res.Header.DisableNormalizing()

	req.Header.SetMethod(aw.AW.Req.Method)
	req.SetRequestURI(aw.AW.Req.Url)
	for key, value := range aw.AW.Req.Header {
		req.Header.Set(key, value)
		aw.AW.Req.HeaderLen += len(key) + len(value) + 4
	}
	if len(req.Header.ContentType()) == 0 {
		req.Header.SetContentType(DefaultContentTypeValue)
		aw.Req.Header[DefaultContentTypeKey] = DefaultContentTypeValue
	}
	req.SetBodyString(aw.AW.Req.Body)
	var err error
	if err = DoRequest(req, res, aw); err != nil {
		aw.Res.ErrReason = err.Error()
		return
	}
	if aw.Req.RedirectsEnabled { // 开启重定向
		redirectsCount := 0
		for res.StatusCode() == http.StatusFound {
			redirectsCount++
			if redirectsCount > maxRedirectsCount {
				aw.AW.Res.ErrReason = http.StatusText(http.StatusTooManyRequests)
				return
			}
			req.URI().UpdateBytes(res.Header.Peek("Location"))
			if err = DoRequest(req, res, aw); err != nil {
				aw.Res.ErrReason = err.Error()
				return
			}
		}
	}
	//
	aw.Result = true
	aw.AW.Res.StatusCode = res.StatusCode()
	res.Header.VisitAll(func(key, value []byte) {
		if v, ok := aw.AW.Res.Header[string(key)]; ok {
			aw.AW.Res.Header[string(key)] = v + ";" + string(value)
		} else {
			aw.AW.Res.Header[string(key)] = string(value)
		}
		aw.AW.Res.HeaderLen += len(key) + len(value) + 4
	})
	aw.AW.Res.Body = util.ByteToString(res.Body())
}

func DoRequest(req *fasthttp.Request, resp *fasthttp.Response, aw *base.AWResult) (err error) {
	if aw.AW.Req.TimeOut.Connect != 0 {
		err = fastHttpClient.DoTimeout(req, resp, time.Duration(aw.AW.Req.TimeOut.Connect))
	} else {
		err = fastHttpClient.Do(req, resp)
	}
	return err
}
