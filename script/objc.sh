#!/bin/bash
# JSON schema files. Figure out the path based on the location of this script.
THIS_SCRIPT=$(readlink -f $0)
# This script's directory
THIS_DIR="$(dirname "${THIS_SCRIPT}")"
# Project directory is parent of this script
PROJECT_DIR="$(dirname "${THIS_DIR}")"
# source generation script is in this script's directory
GEN_SRC_SCRIPT="${THIS_DIR}/index.sh"
# Objective C source location
OBJC_SRC_DIR="${PROJECT_DIR}/objc_build"
# ensure the source directory exists
mkdir -p ${OBJC_SRC_DIR}
# remove any old source code laying around that may not get overwritten
if [ -d ${OBJC_SRC_DIR} -a -w ${OBJC_SRC_DIR} ]; then
  find ${OBJC_SRC_DIR} -type f -exec rm {} \;
fi
# generate objc source
${GEN_SRC_SCRIPT} objc ${OBJC_SRC_DIR}
# generate protobuf source
protoc "${PROJECT_DIR}/proto/Msg.proto" --proto_path="${PROJECT_DIR}" --objc_out="${OBJC_SRC_DIR}"
protoc "${PROJECT_DIR}/proto/Feature.proto" --proto_path="${PROJECT_DIR}" --objc_out="${OBJC_SRC_DIR}"
