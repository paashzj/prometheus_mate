# prometheus_mate
## version
- prometheus 2.36.1
## 环境变量举例
### 添加coredns，只保留一个指标
```bash
-e COREDNS_TYPE=static -e COREDNS_HOSTS=127.0.0.1 -e COREDNS_KEEP_METRICS=coredns_dns_requests_total
```
## motivation
Prometheus原生只能通过修改配置文件的方式来重新加载，有的时候非常不方便，本工程通过开发`HTTP` API，来简化上述的流程。
## fast config from env
### env
#### ${prefix}_TYPE
#### ${prefix}_PORT
config the port, some famous service has default port, see below
#### ${prefix}_METRICS_PATH
config the metrics path
#### ${prefix}_KEEP_METRICS
metrics white lists
#### ${prefix}_SSL
#### ${prefix}_HOSTS
type为static的时候使用，hosts
#### ${prefix}_DOMAINS
type为dns的时候使用，域名
### support famous port
- ZOOKEEPER
- BOOKKEEPER
- PULSAR
- PULSAR_PROXY
- MYSQL
- REDIS
- COREDNS
## 选型框架
GoFrame

## test command
- [Test Command Doc](test_command.md)
