#!/bin/bash

# instructions taken from here:
# http://golangcookbook.com/chapters/running/cross-compiling/

cmd="arch=\$GOOS-\$GOARCH; mkdir builds/\$arch;go build -o builds/\$arch/stake -x stake.go"
GOOS=darwin GOARCH=amd64       eval "$cmd"
GOOS=linux  GOARCH=amd64       eval "$cmd"
GOOS=linux  GOARCH=arm GOARM=7 eval "$cmd"