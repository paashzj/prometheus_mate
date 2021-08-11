package model

type SdConfigs struct {

	// sd type
	SdType string `json:"sd_type,omitempty"`

	StaticConfigs []StaticSdConfig `json:"static_configs,omitempty"`

	FileSdConfigs []FileSdConfig `json:"file_sd_configs,omitempty"`

	DnsSdConfigs []DnsSdConfig `json:"dns_sd_configs,omitempty"`
}
