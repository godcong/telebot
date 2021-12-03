#!/usr/bin/env bash

echo "Clone required repositories"
git clone https://github.com/godcong/go-tdlib.git
git clone https://github.com/tdlib/td.git

echo "Copying files..."
cp -f ./backup/client/tdjson_dynamic.go ./go-tdlib/client/
cp -f ./backup/client/tdlib.go ./go-tdlib/client/

echo "Remove go.mod"
rm ./go-tdlib/go.mod