#!/bin/bash

set -x

pid=$(ps aux | grep "mytholojam-linux" | awk '{print $2}')
kill -9 $pid
cd ~/bin/mytholojam
mv tmp/* .
nohup ./mytholojam-linux 6984 &
sleep 2
disown
