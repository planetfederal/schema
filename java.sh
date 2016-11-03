#!/bin/sh
node index.js java
mkdir -p src/main/java/com/boundlessgeo/spatialconnect/schema
mv SCCommand.java src/main/java/com/boundlessgeo/spatialconnect/schema/
protoc SCMessage.proto --java_out="./"

echo "package com.boundlessgeo.spatialconnect.schema;" > tmp
cat SCMessageOuterClass.java >> tmp
rm SCMessageOuterClass.java
mv tmp src/main/java/com/boundlessgeo/spatialconnect/schema/SCMessageOuterClass.java

./gradlew build
./gradlew install
./gradlew uploadArchives

rm -rf src
