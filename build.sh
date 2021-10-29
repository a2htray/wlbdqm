#!/bin/bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/wlbdqm_linux_amd64 ./cmd/main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./bin/wlbdqm_darwin_amd64 ./cmd/main.go
