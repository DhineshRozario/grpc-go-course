#!/bin/bash

protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.

protoc squareroot/protocolbuffer/calculator.proto --go_out=plugins=grpc:.

protoc primenumber/protocolbuffer/primenumber.proto  --go_out=plugins=grpc:.

protoc average/protocolbuffer/average.proto  --go_out=plugins=grpc:.