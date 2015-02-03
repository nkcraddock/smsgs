#!/bin/bash
ulimit -n 1024
chown -R rabbitmq:rabbitmq /data
exec rabbitmq-server $@ &
sleep 10
  /smsgs/bin/dispatcher & 
  /smsgs/bin/webapi -p 80
