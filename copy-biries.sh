#!/usr/bin/env bash

export TAG="v1.2.8"
docker pull footprintai/grandturks-smartcity-client:$TAG
docker create -it --name smartcityclient footprintai/grandturks-smartcity-client:$TAG /bin/bash
docker cp smartcityclient:/out/. release
docker rm -f smartcityclient
