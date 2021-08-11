package generate

import (
	"encoding/json"
	"github.com/gogf/gf/util/gconv"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"prometheus_mate/pkg/config"
	"prometheus_mate/pkg/constant"
	"prometheus_mate/pkg/model"
	"prometheus_mate/pkg/path"
	"strings"
)

// All write method must add \n on end

func ConvProm() string {
	var sb strings.Builder
	sb.WriteString("global:\n")
	if config.GLOBAL_SCRAPE_INTERVAL == "" {
		sb.WriteString("  scrape_interval: 15s\n")
	} else {
		sb.WriteString("  scrape_interval: " + config.GLOBAL_SCRAPE_INTERVAL + "\n")
	}
	if config.GLOBAL_EVALUATION_INTERVAL == "" {
		sb.WriteString("  evaluation_interval: 15s\n")
	} else {
		sb.WriteString("  evaluation_interval: " + config.GLOBAL_EVALUATION_INTERVAL + "\n")
	}
	if config.GLOBAL_SCRAPE_TIMEOUT == "" {
		sb.WriteString("  scrape_timeout: 10s\n")
	} else {
		sb.WriteString("  scrape_timeout: " + config.GLOBAL_SCRAPE_TIMEOUT + "\n")
	}
	sb.WriteString("scrape_configs:\n")
	// iterate the job
	filepath.Walk(path.PromJobs, func(path string, info fs.FileInfo, err error) error {
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

func convJob(req model.CreateJobReq) string {
	var sb strings.Builder
	// write a job
	_, _ = sb.WriteString("  - job_name: \"" + req.Job + "\"\n")
	if req.TlsConfig != nil {
		_, _ = sb.WriteString("    tls_config:\n")
		_, _ = sb.WriteString("        ca_file: " + path.PromCaPath + "\n")
		_, _ = sb.WriteString("        cert_file: " + path.PromCertPath + "\n")
		_, _ = sb.WriteString("        key_file: " + path.PromKeyPath + "\n")
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
		_, _ = sb.WriteString("    metrics_path: " + req.MetricPath + "\n")
	}
	return sb.String()
}

func convStaticJobs(configs []model.StaticSdConfig) string {
	var sb strings.Builder
	sb.WriteString("     static_configs:\n")
	var aux strings.Builder
	for _, config := range configs {
		aux.WriteString("\"" + config.Name + "\"" + ",")
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
