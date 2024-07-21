package language

import (
	"bytes"
	"embed"
	"text/template"

	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

//go:embed "template"
var mt embed.FS

type ClientCode struct {
	req     *pluginpb.CodeGeneratorRequest
	file    *descriptorpb.FileDescriptorProto
	service *descriptorpb.ServiceDescriptorProto
	ns      *Namespace
	embed   embed.FS
}

func NewClientCode(req *pluginpb.CodeGeneratorRequest, file *descriptorpb.FileDescriptorProto, service *descriptorpb.ServiceDescriptorProto, ns *Namespace) *ClientCode {
	return &ClientCode{
		req:     req,
		file:    file,
		service: service,
		ns:      ns,
		embed:   mt,
	}
}

func (c ClientCode) Body() (string, error) {
	out := bytes.NewBuffer(nil)
	data := struct {
		Namespace *Namespace
		File      *descriptorpb.FileDescriptorProto
		Service   *descriptorpb.ServiceDescriptorProto
	}{
		Namespace: c.ns,
		File:      c.file,
		Service:   c.service,
	}
	tpl := template.New("client.tpl").Funcs(template.FuncMap{
		"client": func(name *string) string {
			return c.ns.p.Identifier(*name, "client")
		},
		"name": func(ns *Namespace, name *string) string {
			return ns.resolve(name)
		},
	})
	t, err := tpl.ParseFS(c.embed, "template/client.tpl")
	if err != nil {
		return "", err
	}
	err = t.Execute(out, data)
	return out.String(), nil
}
