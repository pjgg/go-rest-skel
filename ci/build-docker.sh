#!/bin/bash -x

source ci/build.env.sh


$DOCKER build -t $DOCKER_IMAGE . 

