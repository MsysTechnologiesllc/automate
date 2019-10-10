#!/bin/bash

printf 'GEN: %s\n' components/automate-grpc/protoc-gen-policy/api/*.proto
protoc --go_out=logtostderr=true,paths=source_relative:/src \
  components/automate-grpc/protoc-gen-policy/api/*.proto

printf 'GEN: %s\n' components/automate-grpc/protoc-gen-policy/iam/*.proto
protoc --go_out=logtostderr=true,paths=source_relative:/src \
  components/automate-grpc/protoc-gen-policy/iam/*.proto

printf 'GEN: %s\n' lib/grpc/debug/debug_api/*.proto
protoc --go_out=plugins=grpc,paths=source_relative:/src \
  --policy_out=logtostderr=true,paths=source_relative:/src \
  lib/grpc/debug/debug_api/*.proto

printf 'GEN: %s\n' components/automate-grpc/protoc-gen-a2-config/api/**/*.proto
protoc --go_out=logtostderr=true,paths=source_relative:/src \
  components/automate-grpc/protoc-gen-a2-config/api/**/*.proto
