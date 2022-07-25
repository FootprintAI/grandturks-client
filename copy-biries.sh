#!/usr/bin/env bash

docker create -it --name smartcityclient footprintai/grandturks-smartcity-client:v1 /bin/bash
docker cp smartcityclient:/out/. release
docker rm -f smartcityclient
