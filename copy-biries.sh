#!/usr/bin/env bash

export TAG="v2.2.1"
docker pull footprintai/grandturks-kafeido-client:$TAG
docker create -it --name kafeidoclient footprintai/grandturks-kafeido-client:$TAG /bin/bash
docker cp kafeidoclient:/out/. kafeido
docker rm -f kafeidoclient
