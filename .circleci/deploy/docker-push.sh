#!/bin/bash

set -e

aws ecr get-login --no-include-email --region $AWS_DEFAULT_REGION | bash

docker tag "statistico-odds-warehouse-console" "$AWS_ECR_ACCOUNT_URL/statistico-odds-warehouse:latest"
docker push "$AWS_ECR_ACCOUNT_URL/statistico-odds-warehouse:latest"
