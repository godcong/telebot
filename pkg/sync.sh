#!/usr/bin/env bash

git clone https://github.com/zelenin/go-tdlib.git
git clone https://github.com/tdlib/td.git

cp -f ./backup/client/tdjson_dynamic.go ./go-tdlib/td/client/
cp -f ./backup/client/tdlib.go ./go-tdlib/td/client/
rm ./go-tdlib/go.mod