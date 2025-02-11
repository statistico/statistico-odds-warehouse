#!/bin/bash

set -e

mkdir -p /tmp/workspace/artifacts

GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build -a -o bootstrap ./lambda/main.go

zip /tmp/workspace/artifacts/statistico-odds-warehouse.zip bootstrap
