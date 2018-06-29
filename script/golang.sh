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
# Project directory is parent of script directory
PROJECT_DIR="$(dirname "${SCRIPT_DIR}")"
# Golang source location
GO_SRC_DIR="${PROJECT_DIR}/go_build"
# ensure the source directory exists
mkdir -p ${GO_SRC_DIR}
# remove any old source code laying around that may not get overwritten
if [ -d ${GO_SRC_DIR} -a -w ${GO_SRC_DIR} ]; then
  find ${GO_SRC_DIR} -type f -exec rm {} \;
fi
# pkg directory
PKG_DIR=$SCRIPT_DIR/../proto/boundlessgeo_schema/
# switch to the script dir in case of running the script ouside of this dir
cd $SCRIPT_DIR
# build protobufs
protoc -I=$PROTO_DIR --go_out=plugins=grpc:$DST_DIR $PKG_DIR/*.proto
#source generation script
GEN_SRC_SCRIPT="${SCRIPT_DIR}/index.sh"

${GEN_SRC_SCRIPT} golang "${DST_DIR}/boundlessgeo_schema"
