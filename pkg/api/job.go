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

package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"net/http"
	"prometheus_mate/pkg/model"
	"prometheus_mate/pkg/prom"
	"prometheus_mate/pkg/service"
)

func AddJob(r *ghttp.Request) {
	var req model.CreateJobReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteStatusExit(http.StatusInternalServerError)
	}
	resp, err := service.AddJob(req)
	if err != nil {
		glog.Info("err occurred ", err)
		r.Response.WriteStatusExit(http.StatusInternalServerError)
	}
	prom.ReloadChannel <- struct{}{}
	r.Response.WriteHeader(http.StatusCreated)
	err = r.Response.WriteJsonExit(&resp)
	if err != nil {
		return
	}
}

func DelJob(r *ghttp.Request) {
	job := gconv.String(r.Get("job"))
	if job == "" {
		glog.Info("job name is empty, do nothing")
		r.Response.WriteStatusExit(http.StatusBadRequest, "job name empty")
	}
	err := service.DelJob(job)
	if err != nil {
		glog.Error("err occurred ", err)
		r.Response.WriteStatusExit(http.StatusInternalServerError)
	}
	prom.ReloadChannel <- struct{}{}
	r.Response.WriteStatusExit(http.StatusNoContent)
}
