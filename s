#!/bin/bash

port=$1

if [ -z "$port" ]; then
  port="8080"
fi

set -x

go build -o bin/mytholojam-server ./server
./bin/mytholojam-server $port
