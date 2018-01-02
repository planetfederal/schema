To add functionality, like new methods or parameter/response types:

1. Modify worm.proto as needed
2. To refresh the go server code, from proto directory:

        protoc -I worm/ worm/worm.proto --go_out=plugins=grpc:worm

3. To refresh the python client code, from proto directory:

        python -m grpc_tools.protoc -I worm/ --python_out=worm/ --grpc_python_out=worm/ worm/worm.proto

Then you'll have to modify your server and client code as well.