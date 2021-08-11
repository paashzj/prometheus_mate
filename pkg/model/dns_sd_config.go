package model

type DnsSdConfig struct {
	Names []string `json:"names,omitempty"`

	Type_ string `json:"type,omitempty"`

	Port int `json:"port,omitempty"`

	RefreshInterval string `json:"refresh_interval,omitempty"`
}
