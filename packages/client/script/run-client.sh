#!/bin/bash

sh stop-client.sh

# 以后台方式启动程序，并且将日志记录到 app.log
nohup ./client >> client.log &
