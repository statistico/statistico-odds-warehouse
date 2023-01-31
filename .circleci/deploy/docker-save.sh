#!/bin/bash

set -e

mkdir -p /tmp/workspace/docker-cache

docker save -o /tmp/workspace/docker-cache/statisticooddswarehouse_console.tar statistico-odds-warehouse-console:latest
