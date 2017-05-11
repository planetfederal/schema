#!/bin/sh
node index.js java
mkdir -p src/main/java/com/boundlessgeo/spatialconnect/schema
mv SCCommand.java src/main/java/com/boundlessgeo/spatialconnect/schema/
protoc APIMessage.proto --java_out="./src/main/java"
protoc Feature.proto --java_out="./src/main/java"

./gradlew build
./gradlew install
./gradlew uploadArchives

#rm -rf src
