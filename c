#!/bin/bash

server=$1

if [[ -z "$server" ]]; then
  server="http://127.0.0.1:8080"
fi

set -x

go build -o bin/mytholojam-cli ./cli
./bin/mytholojam-cli $server
