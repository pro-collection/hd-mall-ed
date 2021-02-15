#!/bin/bash


CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./admin ./main.go

scp ./admin yanle-tx:/root/app/hd_mell/admin
scp ../common/config/appDev.ini yanle-tx:/root/app/hd_mell/packages/common/config/appDev.ini

