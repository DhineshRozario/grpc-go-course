#!/bin/bash

# protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.

protoc calculator/protocolbuffer/calculator.proto --go_out=plugins=grpc:.