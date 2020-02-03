#!/bin/bash

module=$(sed -n 's/^module \(.*\)/\1/p' go.mod)

docker run \
  --rm \
  -p 6060:6060 \
  -v $PWD:/go/src/goevent \
  golang:1.13 \
  bash -c "cd src && go get golang.org/x/tools/cmd/godoc && echo http://localhost:6060/pkg/goevent && /go/bin/godoc -http=:6060"