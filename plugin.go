// Copyright (c) Huawei Technologies Co., Ltd. 2020-2020. All rights reserved.
// Description:
// Author: l30002214
// Create: 2020/8/19

// Package main 
package main

import "engine-go/base"

// 插件系统的解析对接接口
type PluginFunc interface {
	SetBase(*base.BaseCase)
	Setup()
	Test()
	TearDown()
}

func NewCaseName() PluginFunc {
	return new(CaseName)
}

type CaseName struct {
	*base.BaseCase
}

func (c *CaseName) SetBase(v *base.BaseCase) {
	c.BaseCase = v
}

func (c *CaseName) Setup() {
	c.Clients[base.Http].Exec(nil)
}

func (c *CaseName) Test() {
	c.Clients[base.Http].Exec(nil)
}

func (c *CaseName) TearDown() {
	c.Clients[base.Http].Exec(nil)
}
