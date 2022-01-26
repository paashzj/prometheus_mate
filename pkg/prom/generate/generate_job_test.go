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

package generate

import (
	"github.com/stretchr/testify/assert"
	"prometheus_mate/pkg/constant"
	"prometheus_mate/pkg/model"
	"testing"
)

func TestConvStaticJob(t *testing.T) {
	staticSdConfig := model.StaticSdConfig{}
	staticSdConfig.Targets = []string{"localhost:9090"}
	result := convStaticJob(staticSdConfig)
	expect := `       - targets: ["localhost:9090"]
`
	assert.Equal(t, expect, result)
}

func TestConvFileJob(t *testing.T) {
	jobReq := model.CreateJobReq{}
	jobReq.Job = "file_job"
	tlsConfig := &model.TlsConfig{}
	tlsConfig.KeyFile = "key.pem"
	tlsConfig.CertFile = "cert.pem"
	tlsConfig.CaFile = "ca.pem"
	jobReq.TlsConfig = tlsConfig
	jobReq.MetricPath = "/scirtem"
	sdConfigs := &model.SdConfigs{}
	sdConfigs.SdType = constant.SdFile
	sdConfigs.FileSdConfigs = make([]model.FileSdConfig, 1)
	fileSdConfig := model.FileSdConfig{}
	fileSdConfig.Files = []string{"file1", "file2"}
	fileSdConfig.RefreshInterval = "10s"
	sdConfigs.FileSdConfigs[0] = fileSdConfig
	jobReq.SdConfigs = sdConfigs
	result := convJob(jobReq)
	expect := `  - job_name: "file_job"
    tls_config:
        ca_file: /tls/ca.pem
        cert_file: /tls/client.pem
        key_file: /tls/client-key.pem
    file_sd_configs:
      - files:
          - "file1"
          - "file2"
        refresh_interval: 10s
    metrics_path: /scirtem
`
	assert.Equal(t, expect, result)
}

func TestConvDnsJob(t *testing.T) {
	dnsSdConfig := model.DnsSdConfig{}
	dnsSdConfig.Names = []string{}
	dnsSdConfig.Type_ = "A"
	dnsSdConfig.Port = 8080
	dnsSdConfig.RefreshInterval = "10s"
	result := convDnsJob(dnsSdConfig)
	expect := `      - names:
        type: "A"
        port: 8080
        refresh_interval: 10s
`
	assert.Equal(t, expect, result)
}

func TestConvKeepMetricsJob(t *testing.T) {
	jobReq := model.CreateJobReq{}
	jobReq.Job = "keep_metrics_job"
	jobReq.KeepMetrics = "request_total|response_total"
	dnsSdConfig := model.DnsSdConfig{}
	dnsSdConfig.Names = []string{}
	dnsSdConfig.Type_ = "A"
	dnsSdConfig.Port = 8080
	dnsSdConfig.RefreshInterval = "10s"
	expect := `  - job_name: "keep_metrics_job"
    dns_sd_configs:
    metric_relabel_configs:
      - source_labels: [__name__]
        regex: (request_total|response_total)
        action: keep
`
	sdConfigs := &model.SdConfigs{}
	sdConfigs.SdType = constant.SdDns
	//sdConfigs.DnsSdConfigs = make([]model.DnsSdConfig, 1)

	//sdConfigs.FileSdConfigs[0] = fileSdConfig
	jobReq.SdConfigs = sdConfigs
	result := convJob(jobReq)
	assert.Equal(t, expect, result)
}
