#!/bin/bash

cd ~/goFirefoxApi

go build

mkdir /temp/goFirefoxApi

cp goFirefoxApi /tmp/goFirefoxApi/

cd /temp/goFirefoxApi

nohup ./goFirefoxApi  >> ~/goLog 2>&1 &

echo "$!" > ~/goPid

echo "GO server started in the background with PID: $!"