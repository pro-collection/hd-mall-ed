#!/bin/bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./client ./../main.go

scp ./client yanle-tx:/root/app/hd_mell/client
scp ./run-client.sh yanle-tx:/root/app/hd_mell/run-client.sh
scp ../common/config/appDev.ini yanle-tx:/root/app/hd_mell/packages/common/config/appDev.ini
