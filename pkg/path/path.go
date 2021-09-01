package path

import (
	"os"
	"path/filepath"
)

var (
	PromHome     = os.Getenv("PROM_HOME")
	PromJobs     = filepath.FromSlash(PromStorage + "/jobs")
	PromConfig   = filepath.FromSlash(PromHome + "/prometheus.yml")
	PromTls      = filepath.FromSlash(PromHome + "/tls")
	PromCaPath   = filepath.FromSlash(PromTls + "/ca.pem")
	PromCertPath = filepath.FromSlash(PromTls + "/client.pem")
	PromKeyPath  = filepath.FromSlash(PromTls + "/client-key.pem")
)

// mate
var (
	PromMatePath      = filepath.FromSlash(PromHome + "/mate")
	PromScripts       = filepath.FromSlash(PromMatePath + "/scripts")
	PromStartScript   = filepath.FromSlash(PromScripts + "/start-prom.sh")
	PromReStartScript = filepath.FromSlash(PromScripts + "/restart-prom.sh")
	PromStorage       = filepath.FromSlash(PromMatePath + "/storage")
)
