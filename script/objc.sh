#!/bin/sh
node script/index.js objc
mkdir -p objc_build
mv Actions.* objc_build/
protoc ./proto/Msg.proto --objc_out="./objc_build"
protoc ./proto/Feature.proto --objc_out="./objc_build"

