package config

import (
	"os"
)

var (
	GlobalScrapeInterval     = os.Getenv("GLOBAL_SCRAPE_INTERVAL")
	GlobalEvaluationInterval = os.Getenv("GLOBAL_EVALUATION_INTERVAL")
	GlobalScrapeTimeout      = os.Getenv("GLOBAL_SCRAPE_TIMEOUT")
)
