// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package prom

import (
	"bufio"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gproc"
	"os"
	"prometheus_mate/pkg/path"
	"prometheus_mate/pkg/prom/generate"
	"prometheus_mate/pkg/util"
)

var ReloadChannel = make(chan struct{})

func Start() {
	startProm()
	go func() {
		for {
			<-ReloadChannel
			startOrReloadProm()
		}
	}()
}

func startOrReloadProm() {
	exists, err := util.ProcessExists(path.PromHome + "/prometheus")
	if err != nil {
		glog.Error("unknown prometheus exists ", err)
		return
	}
	if exists {
		restartProm()
	} else {
		startProm()
	}
}

func startProm() {
	err := generatePromFile()
	if err != nil {
		glog.Error("generate prom config file failed ", err)
		return
	}
	err = gproc.ShellRun("bash -x " + path.PromStartScript)
	glog.Error("run start prom scripts failed ", err)
}

func restartProm() {
	err := generatePromFile()
	if err != nil {
		glog.Error("generate prom config file failed ", err)
		return
	}
	err = gproc.ShellRun("bash -x " + path.PromReStartScript)
	glog.Error("run restart prom scripts failed ", err)
}

func generatePromFile() (err error) {
	file, err := gfile.OpenWithFlag(path.PromConfig, os.O_RDWR|os.O_CREATE|os.O_TRUNC)
	if err != nil {
		glog.Error("open file failed", err)
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(generate.ConvProm())
	if err != nil {
		glog.Error("write string failed", err)
		return
	}
	err = writer.Flush()
	return
}
