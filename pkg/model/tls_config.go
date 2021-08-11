package model

type TlsConfig struct {

	// ca file path
	CaFile string `json:"ca_file,omitempty"`

	// cert file path
	CertFile string `json:"cert_file,omitempty"`

	// key file path
	KeyFile string `json:"key_file,omitempty"`
}
