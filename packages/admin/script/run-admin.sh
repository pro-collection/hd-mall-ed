#!/bin/bash

#sh /root/app/hd_mell/stop-admin.sh

# 以后台方式启动程序，并且将日志记录到 app.log
nohup /root/app/hd_mell/admin >> /root/app/hd_mell/admin.log &
