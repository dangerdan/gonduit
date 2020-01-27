#!/bin/bash
set -euo pipefail
IFS=$'\n\t'

rm -r /home/drydock/go/src

mkdir -p /home/drydock/go/src/github.com/dangerdan
ln -sf $(pwd) /home/drydock/go/src/github.com/dangerdan/
cd /home/drydock/go/src/github.com/dangerdan/gonduit

glide install
go build $(glide novendor)
go test $(glide novendor)
