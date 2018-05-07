#!/bin/bash
# present working dir you are running the command from
PWD=`pwd`
# script dir
if [ `uname` == "Linux" ]; then
    SCRIPT_DIR=$(dirname $(readlink -f $0))
else
    SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
fi
# destination dir
DST_DIR=$SCRIPT_DIR/../pkg/
# source directory for proto
PROTO_DIR=$SCRIPT_DIR/../proto
# pkg directory
PKG_DIR=$SCRIPT_DIR/../proto/boundlessgeo_schema/
# switch to the script dir in case of running the script ouside of this dir
cd $SCRIPT_DIR
# build protobufs
protoc -I=$PROTO_DIR --go_out=plugins=grpc:$DST_DIR $PKG_DIR/*.proto

