#!/bin/bash

rm -rf pkg/
mkdir pkg

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build main.go
zip main.zip main

mv main pkg/
mv main.zip pkg/