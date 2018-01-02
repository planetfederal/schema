.PHONY: compile

worm.pb.go: worm.proto
	protoc --go_out=. worm.proto

compile: worm.pb.go
