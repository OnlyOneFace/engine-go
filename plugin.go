// Copyright (c) Huawei Technologies Co., Ltd. 2020-2020. All rights reserved.
// Description:
// Author: l30002214
// Create: 2020/8/19

// Package main 
package main

import "engine-go/base"

func NewCaseName() base.PluginFunc {
	return new(CaseName)
}

type CaseName struct {
	*base.BaseCase
}

func (c *CaseName) SetBase(v *base.BaseCase) {
	c.BaseCase = v
}

func (c *CaseName) Setup() {
	aw := base.NewAWResult()
	client := c.Clients[base.Http]
	client.Exec(aw)
	client.Info("setup")
	aw.Release()
}

func (c *CaseName) Test() {
	aw := base.NewAWResult()
	c.Clients[base.Http].Exec(aw)
	aw.Release()
}

func (c *CaseName) TearDown() {
	aw := base.NewAWResult()
	c.Clients[base.Http].Exec(aw)
	aw.Release()
}
