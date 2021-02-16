#!/bin/bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./client ./../main.go

scp ./client yanle-tx:/root/app/hd_mell/client
scp ./show-log-client.sh yanle-tx:/root/app/hd_mell/show-log-client.sh
scp ./stop-client.sh yanle-tx:/root/app/hd_mell/stop-client.sh
scp ./run-client.sh yanle-tx:/root/app/hd_mell/run-client.sh
scp ../../common/config/appDev.ini yanle-tx:/root/app/hd_mell/packages/common/config/appDev.ini
