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
DST_DIR=$SCRIPT_DIR/../pkg/schema/
# source directory for proto
SRC_DIR=$SCRIPT_DIR/../proto
# switch to the script dir in case of running the script ouside of this dir
cd $SCRIPT_DIR
# build protobufs
protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/Metadata.proto
protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/Command.proto $SRC_DIR/Metadata.proto
protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/Event.proto $SRC_DIR/Metadata.proto
protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/Msg.proto
protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/Feature.proto
protoc -I=$SRC_DIR --go_out=plugins=grpc:$DST_DIR $SRC_DIR/worm.proto \
$SRC_DIR/Command.proto $SRC_DIR/Event.proto
