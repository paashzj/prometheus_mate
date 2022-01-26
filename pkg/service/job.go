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

package service

import (
	"encoding/json"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
	"os"
	"path/filepath"
	"prometheus_mate/pkg/model"
	"prometheus_mate/pkg/path"
)

func AddJob(req model.CreateJobReq) (resp model.CreateJobResp, err error) {
	resp = model.CreateJobResp{}
	resp.Job = req.Job
	file, err := gfile.OpenWithFlag(filepath.FromSlash(path.PromJobs+"/job-"+req.Job+"-v1.json"), os.O_RDWR|os.O_CREATE)
	if err != nil {
		return
	}
	defer file.Close()
	bytes, err := json.Marshal(req)
	if err != nil {
		return
	}
	_, err = file.Write(bytes)
	if err != nil {
		return
	}
	return
}

func DelJob(job string) error {
	glog.Info("begin to delete job ", job)
	err := gfile.Remove(filepath.FromSlash(path.PromJobs + "/job-" + job + "-v1.json"))
	return err
}
