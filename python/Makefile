.PHONY: compile

worm:
	python -m grpc_tools.protoc -I ../proto/ --python_out=boundlessgeo_schema/ --grpc_python_out=boundlessgeo_schema/ ../proto/worm.proto

compile: worm
