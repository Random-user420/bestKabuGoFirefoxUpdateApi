#!/bin/bash

cd ~/bestKabuGoFirefoxUpdateApi

go build

mkdir /temp/bestKabuGoFirefoxUpdateApi

cp goFirefoxApi /tmp/bestKabuGoFirefoxUpdateApi/

cd /temp/bestKabuGoFirefoxUpdateApi

nohup ./goFirefoxApi  >> ~/goLog 2>&1 &

echo "$!" > ~/goPid

echo "GO server started in the background with PID: $!"
