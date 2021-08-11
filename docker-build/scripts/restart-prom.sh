#!/bin/bash

pid=`ps -ef|grep $PROM_HOME/prometheus|grep -v grep|awk '{print $2}'`
kill -HUP $pid