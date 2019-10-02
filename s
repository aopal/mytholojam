#!/bin/bash -x

go build -o bin/mytholojam-server ./server
./bin/mytholojam-server
