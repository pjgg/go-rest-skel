#!/bin/bash

source ci/build.env.sh

echo "VERSION: $VERSION"
echo "GIT_TAG: $GIT_TAG"
echo "GIT_BRANCH: $GIT_BRANCH"

GO_VERSION="1.7"



echo "SRC_PATH="${SRC_PATH}

$DOCKER run --rm -e SRC_PATH=${SRC_PATH} -v $(pwd):/go/${SRC_PATH} golang:$GO_VERSION /go/$SRC_PATH/ci/build-go.sh