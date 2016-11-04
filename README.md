#spatialconnect-schema

## Java
This will put a spatialconnect-schema jar in your local maven repo

```
sh java.sh
```

## Objective-C
### Protocol Buffers

```
protoc SCMessage.proto --objc_out="./"
```

### Actions
To build the actions run
```
node index.js objc
```
It will produce Commands.h.
