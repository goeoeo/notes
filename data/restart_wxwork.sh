#!/usr/bin/env bash

pid=`ps aux |grep WXWork|awk 'NR==2 {print $2}'`
echo $pid

kill -9 $pid

nohup  /opt/apps/com.qq.weixin.work.deepin/files/run.sh 2>&1 &