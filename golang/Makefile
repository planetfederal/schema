.PHONY: compile

PROTOC := $(shell which protoc)
# If protoc isn't on the path, set it to a target that's never up to date, so
# the install command always runs.
ifeq ($(PROTOC),)
    PROTOC = must-rebuild
endif

# Figure out which machine we're running on.
UNAME := $(shell uname)

$(PROTOC):
# Run the right installation command for the operating system.
ifeq ($(UNAME), Darwin)
	brew install protobuf
endif
ifeq ($(UNAME), Linux)
	sudo apt-get install protobuf-compiler
endif

worm.pb.go: ../proto/worm.proto
	protoc --go_out=plugins=grpc:. --proto_path=../proto ../proto/worm.proto

Event.pb.go: ../proto/Event.proto
	protoc --go_out=. --proto_path=../proto ../proto/Event.proto

Command.pb.go: ../proto/Command.proto
	protoc --go_out=. --proto_path=../proto ../proto/Command.proto


compile: worm.pb.go Event.pb.go Command.pb.go
