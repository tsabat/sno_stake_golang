#!/bin/bash

# instructions taken from here:
# http://golangcookbook.com/chapters/running/cross-compiling/

function compile() {
  go build -o "stake-$GOOS-$GOARCH" -x stake.go
}


export GOOS=darwin
export GOARCH=amd64
compile

export GOOS=linux
export GOARCH=amd64
compile

export GOOS=linux
export GOARCH=arm
export GOARM=7
compile