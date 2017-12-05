#!/bin/sh
# JSON schema files. Figure out the path based on the location of this script.
THIS_SCRIPT=$(readlink -f $0)
# This script's directory
THIS_DIR="$(dirname "${THIS_SCRIPT}")"
# Project directory is parent of this script
PROJECT_DIR="$(dirname "${THIS_DIR}")"
# source generation script is in this script's directory
GEN_SRC_SCRIPT="${THIS_DIR}/index.sh"
# Java project root is in PROJECT_DIR/java
JAVA_ROOT_DIR="${PROJECT_DIR}/java"
# Java generated source directory is "src/main/java/com/boundlessgeo/schema" under JAVA_ROOT_DIR
JAVA_SRC_DIR="${JAVA_ROOT_DIR}/src/main/java/com/boundlessgeo/schema"
# Protobuf source directory is "src/main/java" under JAVA_ROOT_DIR
PROTOBUF_SRC_DIR="${JAVA_ROOT_DIR}/src/main/java"
# ensure the source directory exists
mkdir -p ${JAVA_SRC_DIR}
# remove any old source code laying around that may not get overwritten
if [ -d ${JAVA_ROOT_DIR}/src -a -w ${JAVA_ROOT_DIR}/src ]; then
  find ${JAVA_ROOT_DIR}/src -type f -exec rm {} \;
fi
# generate java source
${GEN_SRC_SCRIPT} java ${JAVA_SRC_DIR}
# generate protobuf source
protoc "${PROJECT_DIR}/proto/Msg.proto" --proto_path="${PROJECT_DIR}" --java_out="${PROTOBUF_SRC_DIR}"
protoc "${PROJECT_DIR}/proto/Feature.proto" --proto_path="${PROJECT_DIR}" --java_out="${PROTOBUF_SRC_DIR}"
# build Java artifacts
# geenrate the pom from the version.json
# get the version from version.json
VERSION="$(grep version "${THIS_DIR}"/version.json | awk '{print $2}'| sed -e 's/\"//g')"
# fill in pom.xml with version
cat "${JAVA_ROOT_DIR}/pom.template" | sed -e 's/\${from\.script}/'${VERSION}'/' > "${JAVA_ROOT_DIR}/pom.xml"
mvn clean install -f "${JAVA_ROOT_DIR}/pom.xml"
