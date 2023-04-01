#!/bin/bash

# build windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/the-number-win.exe .

# build linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/the-number-linux .

# build mac
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/the-number-mac .