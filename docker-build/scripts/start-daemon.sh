#!/bin/bash

# make directory
mkdir $PROM_HOME/mate/storage
mkdir $PROM_HOME/mate/storage/jobs

nohup $PROM_HOME/mate/prom_mate >>$PROM_HOME/prom_mate.log 2>>$PROM_HOME/prom_mate_error.log &