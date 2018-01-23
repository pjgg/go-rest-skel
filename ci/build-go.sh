#!/bin/sh

go get -u -v github.com/kardianos/govendor


echo $SRC_PATH
cd ${GOPATH}/$SRC_PATH
pwd
govendor sync

for GOOS in linux; do
  for GOARCH in amd64; do
    echo "Building $GOOS-$GOARCH"
    export GOOS=$GOOS
    export GOARCH=$GOARCH
    go build -ldflags "-w -s" -o bin/app-$GOOS-$GOARCH
  done
done