# Domain service, connected to postgres, with GRPC server

## Update protobuf and grpc code
    $ protoc --go_out=plugins=grpc:pb service.proto
