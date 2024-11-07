#!/bin/bash

cd /home/lilith/bestKabuGoFirefoxUpdateApi

go build

mkdir /tmp/bestKabuGoFirefoxUpdateApi

cp goFirefoxApi /tmp/bestKabuGoFirefoxUpdateApi/

cd /temp/bestKabuGoFirefoxUpdateApi

nohup ./goFirefoxApi  >> /home/lilith/goLog 2>&1 &

echo "$!" > ~/goPid

echo "GO server started in the background with PID: $!"
