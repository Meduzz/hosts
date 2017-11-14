#!/bin/bash
GOOS=darwin GOARCH=amd64 go build -o hosts_osx_x86_64
GOOS=linux GOARCH=amd64 go build -o hosts_linux_x86_64
GOOS=linux GOARCH=arm GOARM=5 go build -o hosts_linux_pi