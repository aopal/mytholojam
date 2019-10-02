#!/bin/bash -x

server="http://127.0.0.1:8080"

go build -o bin/mytholojam-cli ./cli
./bin/mytholojam-cli $server
