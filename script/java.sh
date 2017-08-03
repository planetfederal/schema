#!/bin/sh
node script/index.js java
mkdir -p ./java/src/main/java/com/boundlessgeo/schema
mv Actions.java ./java/src/main/java/com/boundlessgeo/schema/
protoc ./proto/Msg.proto --java_out="./java/src/main/java"
protoc ./proto/Feature.proto --java_out="./java/src/main/java"

cd java
mvn install
cd ..
