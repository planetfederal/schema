# Schema

## Version 0.12

The Schema repository holds the messaging protocols for data transmission between projects. The commands below will create Protobuf classes for use in iOS, Android, Clojure, and Java. Javascript will only use the Actions.json since the contents of the file is already native to the language. By convention, Actions can be Queries or Commands in a [CQRS](https://martinfowler.com/bliki/CQRS.html) architecture. For mobile `v1/CONFIG_FULL` is a query in the Action field, and `v1/REGISTER_DEVICES` is a Command. The destination/handler of the messages are aware of what is a Query and what is an Command from the producer.

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
