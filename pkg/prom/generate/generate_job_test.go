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
