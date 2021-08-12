# prometheus_mate
## motivation
Prometheus原生只能通过修改配置文件的方式来重新加载，有的时候非常不方便，本工程通过开发`HTTP` API，来简化上述的流程。
## fast config from env
### env
#### ${prefix}_TYPE
#### ${prefix}_SSL
#### ${prefix}_HOSTS
type为static的时候使用，hosts
#### ${prefix}_DOMAINS
type为dns的时候使用，域名
#### ${prefix}_METRICS_PATH
### support env prefix
- ZOOKEEPER
- BOOKKEEPER
- PULSAR
- PULSAR_PROXY
- MYSQL
- REDIS
## 选型框架
GoFrame

## test command
- [Test Command Doc](test_command.md)