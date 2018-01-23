#!/bin/bash -x

source ci/build.env.sh

export CONTAINER_NAME=${JOB_NAME}-${BUILD_NUMBER}
$DOCKER rm -f $CONTAINER_NAME

$DOCKER run -d --name $CONTAINER_NAME \
	-v $(pwd)/config.yaml:/config.yaml \
	$DOCKER_IMAGE /app

$DOCKER run -t --rm \
	-v $(pwd):/go/${SRC_PATH} \
	--link DOCKER_IMAGE:person \
	golang:1.7 go test -v -tags=acceptance /go/${SRC_PATH}/acceptance

$DOCKER rm -f $CONTAINER_NAME