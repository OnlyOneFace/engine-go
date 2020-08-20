// Copyright (c) Huawei Technologies Co., Ltd. 2020-2020. All rights reserved.
// Description:
// Author: l30002214
// Create: 2020/8/18

// Package engine_go 
package main

import (
	"log"
	"plugin"

	"engine-go/base"
	"engine-go/impl/client"
	"engine-go/impl/logger"
)

func main() {
	client.Register()
	pluginFile, err := plugin.Open("caseName.so")
	if err != nil {
		log.Fatal(err)
	}
	var symbol plugin.Symbol
	if symbol, err = pluginFile.Lookup("NewCaseName"); err != nil {
		log.Fatal(err)
	}
	cn := symbol.(func() base.PluginFunc)()
	// 日志设置
	for _, value := range base.Clients {
		value.SetLogger(logger.NewZapLogger("Info", true))
	}
	// basecase生成
	var basins = &base.BaseCase{
		PodId:    0,
		PodCount: 0,
		Clients:  base.Clients,
	}
	cn.SetBase(basins)

	cn.Setup()

	cn.Test()

	cn.TearDown()
}
