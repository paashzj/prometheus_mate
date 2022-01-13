#!/bin/bash

# make directory
mkdir -p $PROM_HOME/logs
mkdir -p $PROM_HOME/mate/storage
mkdir -p $PROM_HOME/mate/storage/jobs

nohup $PROM_HOME/mate/prom_mate >>$PROM_HOME/logs/prom_mate.stdout.log 2>>$PROM_HOME/logs/prom_mate.stderr.log &

