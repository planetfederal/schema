#!/bin/bash
# script dir
if [ `uname` == "Linux" ]; then
    SCRIPT_DIR=$(dirname $(readlink -f $0))
else
    SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
fi
# Project directory is parent of script directory
PROJECT_DIR="$(dirname "${SCRIPT_DIR}")"
# source generation script is in this script's directory
GEN_SRC_SCRIPT="${SCRIPT_DIR}/index.sh"
# Objective C source location
OBJC_SRC_DIR="${PROJECT_DIR}/objc_build"
# ensure the source directory exists
mkdir -p ${OBJC_SRC_DIR}
# remove any old source code laying around that may not get overwritten
if [ -d ${OBJC_SRC_DIR} -a -w ${OBJC_SRC_DIR} ]; then
  find ${OBJC_SRC_DIR} -type f -exec rm {} \;
fi
# Protobuf root
PKG_DIR="${PROJECT_DIR}/proto/boundlessgeo_schema"
# source directory for proto
PROTO_DIR="${PROJECT_DIR}/proto/"
# generate objc source
${GEN_SRC_SCRIPT} objc ${OBJC_SRC_DIR}
# generate GRPC not now
#protoc "${PKG_DIR}/worm.proto" "${PKG_DIR}/Command.proto" "${PKG_DIR}/Event.proto" "${PKG_DIR}/Metadata.proto" \
#--proto_path="${PROTO_DIR}" --objc_out=plugins=grpc:"${OBJC_SRC_DIR}"
# generate protobuf source
protoc "${PKG_DIR}/Metadata.proto" --proto_path="${PROTO_DIR}" --objc_out="${OBJC_SRC_DIR}"
protoc "${PKG_DIR}/Command.proto" "${PKG_DIR}/Metadata.proto" --proto_path="${PROTO_DIR}" --objc_out="${OBJC_SRC_DIR}"
protoc "${PKG_DIR}/Event.proto" "${PKG_DIR}/Metadata.proto" --proto_path="${PROTO_DIR}" --objc_out="${OBJC_SRC_DIR}"
protoc "${PKG_DIR}/Msg.proto" --proto_path="${PROTO_DIR}" --objc_out="${OBJC_SRC_DIR}"
protoc "${PKG_DIR}/Feature.proto" --proto_path="${PROTO_DIR}" --objc_out="${OBJC_SRC_DIR}"
