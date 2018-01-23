#!/bin/bash -x

source ci/build.env.sh


go get -u -v github.com/kardianos/govendor

cd ${GOPATH}/$SRC_PATH

govendor sync

go get github.com/go-swagger/go-swagger/cmd/swagger
swagger generate spec -o swagger.json

ls
