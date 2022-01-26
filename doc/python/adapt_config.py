#
# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.
#

import json
import os
import sys

PROM_DIR = sys.argv[1]
PROM_CONFIG = PROM_DIR + "/" + "prometheus.yml"


def common_static(service, env_service_prefix, ssl_enable):
    print("common static " + service + " " + env_service_prefix)
    with open(PROM_CONFIG, "a") as file:
        file.write('  - job_name: "' + service + '"')
        file.write("\n")
        if ssl_enable:
            file.write('    tls_config:')
            file.write('        ca_file: ' + PROM_DIR + '/tls/ca.pem')
            file.write('        cert_file: ' + PROM_DIR + '/tls/client.pem')
            file.write('        key_file: ' + PROM_DIR + '/tls/client-key.pem')
        file.write('    file_sd_configs:')
        file.write("\n")
        file.write('      - files:')
        file.write("\n")
        file.write('          - "/opt/sh/prometheus/' + service + '.json"')
        file.write("\n")
        file.write('        refresh_interval: 10s')
        file.write("\n")
        dns_metric_path = os.getenv(env_service_prefix + "_METRICS_PATH")
        if dns_metric_path is not None:
            file.write('    metrics_path: ' + dns_metric_path)
            file.write("\n")
    hosts = os.getenv(env_service_prefix + "_HOSTS")
    host_array = hosts.split(',')
    one_target = {"targets": host_array}
    targets = [one_target]
    with open(PROM_DIR + "/" + service + ".json", "w") as file:
        file.write(json.dumps(targets))


def common_dns(service, env_service_prefix, port, ssl_enable):
    print("common dns " + service + " " + env_service_prefix + " " + port)
    with open(PROM_CONFIG, "a") as file:
        file.write('  - job_name: "' + service + '"')
        file.write("\n")
        if ssl_enable:
            file.write('    tls_config:')
            file.write('        ca_file: ' + PROM_DIR + '/tls/ca.pem')
            file.write('        cert_file: ' + PROM_DIR + '/tls/client.pem')
            file.write('        key_file: ' + PROM_DIR + '/tls/client-key.pem')
        file.write('    dns_sd_configs:')
        file.write("\n")
        file.write('      - names:')
        file.write("\n")
        pulsar_hosts = os.getenv(env_service_prefix + "_DOMAINS")
        pulsar_host_array = pulsar_hosts.split(',')
        for host in pulsar_host_array:
            file.write('          - "' + host + '"')
            file.write("\n")
        file.write('        type: "A"')
        file.write("\n")
        file.write('        port: ' + port)
        file.write("\n")
        file.write('        refresh_interval: 10s')
        file.write("\n")
        dns_metric_path = os.getenv(env_service_prefix + "_METRICS_PATH")
        if dns_metric_path is not None:
            file.write('    metrics_path: ' + dns_metric_path)
            file.write("\n")


def common(service, env_service_prefix, port):
    grab_type = os.getenv(env_service_prefix + "_TYPE")
    if grab_type is None:
        return
    ssl_enable = os.getenv(env_service_prefix + "_SSL") is not None
    if grab_type == 'dns':
        common_dns(service, env_service_prefix, port, ssl_enable)
    else:
        common_static(service, env_service_prefix, ssl_enable)


common('zookeeper', "ZOOKEEPER", "7000")
common('bookkeeper', "BOOKKEEPER", "8080")
common('pulsar', "PULSAR", "8080")
common('pulsar_proxy', "PULSAR_PROXY", "8080")
common('mysql', "MYSQL", "9104")
common('redis', "REDIS", "9121")