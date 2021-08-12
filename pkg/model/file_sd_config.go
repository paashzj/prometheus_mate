package model

type FileSdConfig struct {
	Files []string `json:"files,omitempty"`

	RefreshInterval string `json:"refresh_interval,omitempty"`
}
