#!/bin/bash

protoc deadline/protocolbuffer/greetwithdeadline.proto --go_out=plugins=grpc: