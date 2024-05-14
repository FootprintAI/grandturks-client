#!/usr/bin/env bash
#
#
mkdir -p go-openapiv2
swagger generate client  -f openapiv2/*.swagger.json --target go-openapiv2
