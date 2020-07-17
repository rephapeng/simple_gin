#!/bin/bash
set -e

ENV=$1

REVISION=$2

#Your private registry
REGISTRY="<your private registry>"
APP_NAME="simple_gin"


#build go as binary
go build .

#build docker
docker build --build-arg STATE_ENV=$ENV  -t ${REGISTRY}${APP_NAME}:${REVISION} .


#push to registry
#docker push ${REGISTRY}${APP_NAME}:${REVISION}



