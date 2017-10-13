#!/bin/bash

# instructions taken from here:
# http://golangcookbook.com/chapters/running/cross-compiling/

cmd="go build -o stake-\$GOOS-\$GOARCH -x stake.go"
GOOS=darwin GOARCH=amd64       eval "$cmd"
GOOS=linux  GOARCH=amd64       eval "$cmd"
GOOS=linux  GOARCH=arm GOARM=7 eval "$cmd"