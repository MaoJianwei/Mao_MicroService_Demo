#!/bin/bash

set -xe

if [ `whoami` = "root" ]
then
    echo "Should run as user mao"
    exit 2
fi

if [ $# -ne 1 ]
then
    echo "Need version for app image"
    exit 1
fi


cd `dirname $0`

cd src

go clean

protoc --go_out=./grpc.maojianwei.com/api/ --go_opt=paths=source_relative \
       --go-grpc_out=./grpc.maojianwei.com/api/ --go-grpc_opt=paths=source_relative mao.proto

go build -v


BIN_NAME=`awk -F '/' 'NR==1{print $2}' go.mod`
mv ${BIN_NAME} ../

cd ../
sudo docker rmi ${BIN_NAME}:$1 || echo ""
sudo docker build --tag ${BIN_NAME}:$1 .
sudo docker images --all --digests
rm ${BIN_NAME}
