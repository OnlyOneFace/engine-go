// Copyright (c) Huawei Technologies Co., Ltd. 2020-2020. All rights reserved.
// Description:
// Author: l30002214
// Create: 2020/8/18

// Package engine_go 
package main

import "engine-go/base"

func main() {
	cn := NewCaseName()
	var basins = &base.BaseCase{}
	cn.SetBase(basins)

	cn.Setup()

	cn.Test()

	cn.TearDown()
}
