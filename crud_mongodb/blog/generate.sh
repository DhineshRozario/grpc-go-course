#!/bin/bash

protoc blog/protocolbuffer/blog.proto --go_out=plugins=grpc:.