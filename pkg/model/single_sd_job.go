package model

import "prometheus_mate/pkg/constant"

type BaseSdJob struct {
	Job        string
	TlsConfig  *TlsConfig
	MetricPath string
}

type SingleSdJob struct {
	BaseSdJob
	SdType         string
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
