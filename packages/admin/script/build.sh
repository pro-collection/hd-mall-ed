#!/bin/bash

echo "build start ..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./admin ./../main.go
echo "build complete ..."

echo "关闭远端服务 ..."
ssh yanle-tx sh /root/app/hd_mell/stop-admin.sh
echo "关闭远端服务成功 ..."

echo "开始上传文件..."
scp ./admin yanle-tx:/root/app/hd_mell/admin
scp ./show-log-admin.sh yanle-tx:/root/app/hd_mell/show-log-admin.sh
scp ./stop-admin.sh yanle-tx:/root/app/hd_mell/stop-admin.sh
scp ./run-admin.sh yanle-tx:/root/app/hd_mell/run-admin.sh
scp ../../common/config/appDev.ini yanle-tx:/root/app/hd_mell/packages/common/config/appDev.ini
echo "上传文件成功..."


rm ./admin
