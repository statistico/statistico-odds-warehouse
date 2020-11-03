#!/bin/bash

set -e

aws ecr get-login --no-include-email --region $AWS_DEFAULT_REGION | bash

docker tag "statisticoddswarehouse_queue" "$AWS_ECR_ACCOUNT_URL/statistico-odds-warehouse:$CIRCLE_SHA1"
docker push "$AWS_ECR_ACCOUNT_URL/statistico-odds-warehouse:$CIRCLE_SHA1"
