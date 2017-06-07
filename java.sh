#!/bin/sh
node index.js java
mkdir -p src/main/java/com/boundlessgeo/schema
mv Actions.java src/main/java/com/boundlessgeo/schema/
protoc Msg.proto --java_out="./src/main/java"
protoc Feature.proto --java_out="./src/main/java"

./gradlew build
./gradlew install
./gradlew uploadArchives

rm -rf src
