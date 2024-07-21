package main

import (
	"io"

	"google.golang.org/protobuf/proto"
	plugin "google.golang.org/protobuf/types/pluginpb"
)

func main() {
	/*
		req, err := readRequest(os.Stdin)
		if err != nil {
			log.Fatalln(err)
		}

		// if err = writeResponse(os.Stdout, php.Generate(req)); err != nil {
		log.Fatalln(err)
		}
	*/
}

func readRequest(in io.Reader) (*plugin.CodeGeneratorRequest, error) {
	data, err := io.ReadAll(in)
	if err != nil {
		return nil, err
	}

	req := new(plugin.CodeGeneratorRequest)
	if err = proto.Unmarshal(data, req); err != nil {
		return nil, err
	}

	return req, nil
}

func writeResponse(out io.Writer, resp *plugin.CodeGeneratorResponse) error {
	data, err := proto.Marshal(resp)
	if err != nil {
		return err
	}
	_, err = out.Write(data)
	return err
}
