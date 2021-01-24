#!/bin/bash

protoc greet/protocolbuffer/greet.proto --go_out=plugins=grpc:.

protoc calculator/protocolbuffer/calculator.proto --go_out=plugins=grpc:.