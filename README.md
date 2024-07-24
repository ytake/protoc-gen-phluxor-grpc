# protoc-gen-phluxor-grpc

This is a protoc plugin that generates Phluxor gRPC services.

## Install

if use go install

```bash
$ go install github.com/ytake/protoc-gen-phluxor-grpc@latest
```

or download [release binary](https://github.com/ytake/protoc-gen-phluxor-grpc/releases)

```bash
$ cp ./protoc-gen-phluxor-grpc /usr/local/bin/
```

## Usage

```bash
$ protoc --php_out=./path/to \
       --phluxor-grpc_out==./path/to \
       --plugin=protoc-gen-grpc=protoc-gen-phluxor-grpc \
       helloworld.proto
```
