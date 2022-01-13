#!/bin/bash

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
