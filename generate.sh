#!/bin/bash

# protoc unary_grpc/greet/greetpb/greet.proto --go_out=plugins=grpc:.

protoc unary_grpc/calculator/protocolbuffer/calculator.proto --go_out=plugins=grpc:.