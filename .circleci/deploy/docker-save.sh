#!/bin/bash

set -e

mkdir -p /tmp/workspace/docker-cache

docker save -o /tmp/workspace/docker-cache/statisticoddswarehouse_queue.tar statisticoddswarehouse_queue:latest
