#!/bin/bash 

source ci/build.env.sh
$GLOUD $DOCKER -- push $DOCKER_IMAGE

