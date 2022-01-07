package model

type CreateJobReq struct {

	// job name
	Job string `json:"job,omitempty"`

	TlsConfig *TlsConfig `json:"tls_config,omitempty"`

	// metric path
	MetricPath string `json:"metric_path,omitempty"`

	SdConfigs *SdConfigs `json:"sd_configs,omitempty"`

	KeepMetrics string `json:"keep_metrics,omitempty"`
}
