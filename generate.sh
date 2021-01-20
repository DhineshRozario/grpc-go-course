#!/bin/bash

# protoc server_streaming_grpc/greet/greetpb/greet.proto --go_out=plugins=grpc:.

protoc server_streaming_grpc/calculator/protocolbuffer/calculator.proto --go_out=plugins=grpc:.