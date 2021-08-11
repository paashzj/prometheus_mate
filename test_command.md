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