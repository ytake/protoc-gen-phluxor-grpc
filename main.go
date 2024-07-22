package main

import (
	"io"
	"log"
	"os"

	"github.com/ytake/protoc-gen-swoole-grpc/language"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	req, err := readRequest(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}
	code := language.NewPHPCode()
	if err = writeResponse(os.Stdout, code.Generate(req)); err != nil {
		log.Fatalln(err)
	}
}

func readRequest(in io.Reader) (*pluginpb.CodeGeneratorRequest, error) {
	data, err := io.ReadAll(in)
	if err != nil {
		return nil, err
	}
	req := new(pluginpb.CodeGeneratorRequest)
	if err = proto.Unmarshal(data, req); err != nil {
		return nil, err
	}
	return req, nil
}

func writeResponse(out io.Writer, resp *pluginpb.CodeGeneratorResponse) error {
	data, err := proto.Marshal(resp)
	if err != nil {
		return err
	}
	_, err = out.Write(data)
	return err
}
