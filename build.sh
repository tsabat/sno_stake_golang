#!/bin/bash

# instructions taken from here:
# http://golangcookbook.com/chapters/running/cross-compiling/

function compile() {
  go build -o "stake-$GOOS-$GOARCH" -x stake.go
}


GOOS=darwin GOARCH=386
compile

GOOS=linux GOARCH=arm GOARM=7
compile

GOOS=linux GOARCH=386
compile