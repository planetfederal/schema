# Schema

## Version 0.12.0

The Schema repository holds the protocol buffer message schemas used to exchange data between systems.
The commands below will create Protobuf classes for use in iOS, Python, and JVM (Clojure, Java, Android) runtimes.
Javascript will use the `actions.json` and `events.json` since the contents of the file is already native to the language.

## install
First you must install the dependencies for the Javascript script that
will build the language-specific classes.

```
cd script/
npm install
cd -
```

## Java
This will put a schema jar in your local maven repo

```
sh script/java.sh
```

## Objective-C
This will build to a folder called `objc_build` in the root dir

```
sh script/objc.sh
```

## Python
This will build a distribution in the `python/dist` directory

```
bash script/python.sh
```

### Actions
To build the actions run
```
node script/index.js objc
```
It will produce Actions.h and Actions.m.
