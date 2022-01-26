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

# make directory
mkdir -p $PROM_HOME/logs
mkdir -p $PROM_HOME/mate/storage
mkdir -p $PROM_HOME/mate/storage/jobs

nohup $PROM_HOME/mate/prom_mate >>$PROM_HOME/logs/prom_mate.stdout.log 2>>$PROM_HOME/logs/prom_mate.stderr.log &

