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
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"prometheus_mate/pkg/config"
	"prometheus_mate/pkg/constant"
	"prometheus_mate/pkg/model"
	"prometheus_mate/pkg/path"
	"strconv"
	"strings"
)

// All write method must add \n on end

func ConvProm() string {
	var sb strings.Builder
	sb.WriteString("global:\n")
	if config.GlobalScrapeInterval == "" {
		sb.WriteString("  scrape_interval: 15s\n")
	} else {
		sb.WriteString("  scrape_interval: " + config.GlobalScrapeInterval + "\n")
	}
	if config.GlobalEvaluationInterval == "" {
		sb.WriteString("  evaluation_interval: 15s\n")
	} else {
		sb.WriteString("  evaluation_interval: " + config.GlobalEvaluationInterval + "\n")
	}
	if config.GlobalScrapeTimeout == "" {
		sb.WriteString("  scrape_timeout: 10s\n")
	} else {
		sb.WriteString("  scrape_timeout: " + config.GlobalScrapeTimeout + "\n")
	}
	sb.WriteString("scrape_configs:\n")
	sb.WriteString(convPromSelfJob())
	sb.WriteString(convJobFromEnv())
	// iterate the job
	_ = filepath.Walk(path.PromJobs, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".json" {
			bytes, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			var req model.CreateJobReq
			err = json.Unmarshal(bytes, &req)
			if err != nil {
				return err
			}
			_, err = sb.WriteString(convJob(req))
			if err != nil {
				return err
			}
		}
		return nil
	})
	return sb.String()
}

func convPromSelfJob() string {
	req := model.SingleSdJob{}
	req.Job = "prometheus"
	req.SdType = constant.SdStatic
	req.StaticSdConfig = model.StaticSdConfig{Targets: []string{"localhost:9090"}}
	return convJob(req.Conv2Req())
}

func convJobFromEnv() string {
	return autoConvJobFromEnv()
}

func autoConvJobFromEnv() string {
	allEnv := os.Environ()
	var sb strings.Builder
	for _, env := range allEnv {
		if strings.HasSuffix(env, "_TYPE") {
			split := strings.Split(env, "_TYPE")
			service := split[0]
			sb.WriteString(convJobFromEnvService(service, service, getServicePort(service)))
		}
	}
	return sb.String()
}

func convJobFromEnvService(service string, env string, port int) string {
	singleSdJob := model.SingleSdJob{}
	singleSdJob.Job = service
	singleSdJob.SdType = os.Getenv(env + "_TYPE")
	if singleSdJob.SdType == "" {
		return ""
	}
	sslEnable := os.Getenv(env + "_SSL")
	if sslEnable != "" {
		singleSdJob.TlsConfig = &model.TlsConfig{}
		singleSdJob.TlsConfig.CaFile = path.PromCaPath
		singleSdJob.TlsConfig.CertFile = path.PromCertPath
		singleSdJob.TlsConfig.KeyFile = path.PromKeyPath
	}
	if singleSdJob.SdType == constant.SdStatic {
		singleSdJob.StaticSdConfig = model.StaticSdConfig{}
		singleSdJob.StaticSdConfig.Targets = strings.Split(os.Getenv(env+"_HOSTS"), ",")
	} else if singleSdJob.SdType == constant.SdDns {
		singleSdJob.DnsSdConfig = model.DnsSdConfig{}
		singleSdJob.DnsSdConfig.Names = strings.Split(os.Getenv(env+"_DOMAINS"), ",")
		singleSdJob.DnsSdConfig.Type_ = "A"
		singleSdJob.DnsSdConfig.Port = port
		singleSdJob.DnsSdConfig.RefreshInterval = "10s"
	}
	singleSdJob.KeepMetrics = os.Getenv(env + "_KEEP_METRICS")
	return convJob(singleSdJob.Conv2Req())
}

func convJob(req model.CreateJobReq) string {
	var sb strings.Builder
	// write a job
	sb.WriteString("  - job_name: \"" + req.Job + "\"\n")
	if req.TlsConfig != nil {
		sb.WriteString("    tls_config:\n")
		sb.WriteString("        ca_file: " + path.PromCaPath + "\n")
		sb.WriteString("        cert_file: " + path.PromCertPath + "\n")
		sb.WriteString("        key_file: " + path.PromKeyPath + "\n")
	}
	sdConfigs := req.SdConfigs
	if sdConfigs.SdType == constant.SdStatic {
		sb.WriteString(convStaticJobs(sdConfigs.StaticConfigs))
	}
	if sdConfigs.SdType == constant.SdFile {
		sb.WriteString(convFileJobs(sdConfigs.FileSdConfigs))
	}
	if sdConfigs.SdType == constant.SdDns {
		sb.WriteString(convDnsJobs(sdConfigs.DnsSdConfigs))
	}
	if req.MetricPath != "" {
		sb.WriteString("    metrics_path: " + req.MetricPath + "\n")
	}
	if req.KeepMetrics != "" {
		sb.WriteString("    metric_relabel_configs:" + "\n")
		sb.WriteString("      - source_labels: [__name__]" + "\n")
		sb.WriteString("        regex: (" + req.KeepMetrics + ")" + "\n")
		sb.WriteString("        action: keep" + "\n")
	}
	return sb.String()
}

func convStaticJobs(configs []model.StaticSdConfig) string {
	var sb strings.Builder
	sb.WriteString("    static_configs:\n")
	for _, config := range configs {
		sb.WriteString(convStaticJob(config))
	}
	return sb.String()
}

func convStaticJob(config model.StaticSdConfig) string {
	var sb strings.Builder
	var aux strings.Builder
	for _, target := range config.Targets {
		aux.WriteString("\"" + target + "\"" + ",")
	}
	sb.WriteString("       - targets: [" + aux.String()[:len(aux.String())-1] + "]\n")
	return sb.String()
}

func convFileJobs(configs []model.FileSdConfig) string {
	var sb strings.Builder
	sb.WriteString("    file_sd_configs:\n")
	for _, config := range configs {
		sb.WriteString(convFileJob(config))
	}
	return sb.String()
}

func convFileJob(config model.FileSdConfig) string {
	var sb strings.Builder
	sb.WriteString("      - files:\n")
	for _, file := range config.Files {
		sb.WriteString("          - \"" + file + "\"\n")
	}
	sb.WriteString("        refresh_interval: 10s\n")
	return sb.String()
}

func convDnsJobs(configs []model.DnsSdConfig) string {
	var sb strings.Builder
	sb.WriteString("    dns_sd_configs:\n")
	for _, config := range configs {
		sb.WriteString(convDnsJob(config))
	}
	return sb.String()
}

func convDnsJob(config model.DnsSdConfig) string {
	var sb strings.Builder
	sb.WriteString("      - names:\n")
	for _, name := range config.Names {
		sb.WriteString("          - \"" + name + "\"\n")
	}
	sb.WriteString("        type: \"" + config.Type_ + "\"\n")
	sb.WriteString("        port: " + gconv.String(config.Port) + "\n")
	sb.WriteString("        refresh_interval: 10s\n")
	return sb.String()
}

func getServicePort(service string) int {
	port := os.Getenv(service + "_PORT")
	if port == "" {
		switch service {
		case "ZOOKEEPER":
			return 7000
		case "PULSAR":
			return 8080
		case "PULSAR_PROXY":
			return 8080
		case "BOOKKEEPER":
			return 8080
		case "MYSQL":
			return 9104
		case "REDIS":
			return 9121
		case "COREDNS":
			return 9153
		case "KUBELET":
			return 10250
		case "KUBE_PROXY":
			return 10249
		default:
			return 8080
		}
	}
	atoi, err := strconv.Atoi(port)
	if err != nil {
		_ = fmt.Errorf("service %s port %s get from env parse int failed, use defualt port. error: %s", service, port, err)
		return 8080
	}
	return atoi
}
