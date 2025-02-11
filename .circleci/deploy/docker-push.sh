#!/bin/bash

set -e

aws ecr get-login-password --region "$AWS_DEFAULT_REGION" | docker login --username AWS --password-stdin "$AWS_ECR_ACCOUNT_URL/statistico-odds-warehouse"

docker tag "statistico-odds-warehouse-console" "$AWS_ECR_ACCOUNT_URL/statistico-odds-warehouse:latest"
docker push "$AWS_ECR_ACCOUNT_URL/statistico-odds-warehouse:latest"
