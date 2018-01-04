.PHONY: compile

worm:
	python -m grpc_tools.protoc -I ../proto/ --python_out=boundlessgeo_schema/ \
	--grpc_python_out=boundlessgeo_schema/ \
	../proto/Command.proto ../proto/Event.proto ../proto/worm.proto


compile: worm
