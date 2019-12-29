#!/bin/sh
protoc --go_out=plugins=grpc:pb service.proto
cp -r pb ../tests/
cp -r pb ../ClientAPI/
