#!/bin/bash

set -e

BIN=display-test

echo build -------------------------------
env GOOS=linux GOARCH=arm64 GOARM=5 go build -o $BIN

echo copy --------------------------------
rsync -avz $BIN erico@192.168.0.35:/tmp

echo run----------------------------------
ssh erico@192.168.0.35 /tmp/$BIN