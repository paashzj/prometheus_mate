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

package model

import "prometheus_mate/pkg/constant"

type BaseSdJob struct {
	Job        string
	TlsConfig  *TlsConfig
	MetricPath string
}

type SingleSdJob struct {
	BaseSdJob
	SdType string
	StaticSdConfig StaticSdConfig
	FileSdConfig   FileSdConfig
	DnsSdConfig    DnsSdConfig
	KeepMetrics    string
}

func (s SingleSdJob) Conv2Req() CreateJobReq {
	req := CreateJobReq{}
	req.Job = s.Job
	req.TlsConfig = s.TlsConfig
	req.MetricPath = s.MetricPath
	req.SdConfigs = &SdConfigs{}
	req.SdConfigs.SdType = s.SdType
	if s.SdType == constant.SdStatic {
		req.SdConfigs.StaticConfigs = make([]StaticSdConfig, 1)
		req.SdConfigs.StaticConfigs[0] = s.StaticSdConfig
	}
	if s.SdType == constant.SdFile {
		req.SdConfigs.FileSdConfigs = make([]FileSdConfig, 1)
		req.SdConfigs.FileSdConfigs[0] = s.FileSdConfig
	}
	if s.SdType == constant.SdDns {
		req.SdConfigs.DnsSdConfigs = make([]DnsSdConfig, 1)
		req.SdConfigs.DnsSdConfigs[0] = s.DnsSdConfig
	}
	req.KeepMetrics = s.KeepMetrics
	return req
}
