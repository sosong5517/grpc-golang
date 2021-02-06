#bin/bash

protoc greet/greetpb/second.proto --go_out = plugins=grpc:. 