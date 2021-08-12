### run example
```bash
docker run -p 9090:9090 ttbb/prometheus:mate
```

### run fast config example
```bash
docker run -e STORAGE_TSDB_RETENTION_TIME=2d -e STORAGE_TSDB_RETENTION_SIZE=7GB -e GLOBAL_SCRAPE_INTERVAL=30s -e GLOBAL_EVALUATION_INTERVAL=30s -e GLOBAL_SCRAPE_TIMEOUT=20s -e ZOOKEEPER_TYPE=static -e ZOOKEEPER_HOSTS=127.0.0.1 -e BOOKKEEPER_TYPE=static -e BOOKKEEPER_HOSTS=127.0.0.2 -e PULSAR_TYPE=dns -e PULSAR_DOMAINS=pulsar.com -p 9090:9090 ttbb/prometheus:mate
```

### add dns config

```bash
curl -XPOST -H 'content-type: application/json;charset=UTF-8' localhost:31001/v1/prometheus/jobs
```

```json
{
  "job": "string",
  "metric_path": "string",
  "sd_configs": {
    "sd_type": "dns_sd",
    "dns_sd_configs": [
      {
        "names": [
          "pulsar"
        ],
        "type": "A",
        "port": 8080,
        "refresh_interval": "10s"
      }
    ]
  }
}
```