#!/bin/bash
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

if [ $STORAGE_TSDB_RETENTION_TIME ]; then
    PROM_PARAM="$PROM_PARAM --storage.tsdb.retention.time=$STORAGE_TSDB_RETENTION_TIME"
else
    PROM_PARAM="$PROM_PARAM --storage.tsdb.retention.time=1d"
fi
if [ $STORAGE_TSDB_RETENTION_SIZE ]; then
    PROM_PARAM="$PROM_PARAM --storage.tsdb.retention.size=$STORAGE_TSDB_RETENTION_SIZE"
else
    PROM_PARAM="$PROM_PARAM --storage.tsdb.retention.size=5GB"
fi
mkdir -p $PROM_HOME/logs
nohup $PROM_HOME/prometheus $PROM_PARAM --config.file=$PROM_HOME/prometheus.yml >>$PROM_HOME/logs/prometheus.stdout.log 2>>$PROM_HOME/logs/prometheus.stderr.log &
