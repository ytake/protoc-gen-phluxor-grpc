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

type BodyRenderer interface {
	Body() (string, error)
}

type value struct {
	Namespace *Namespace
	File      *descriptorpb.FileDescriptorProto
	Service   *descriptorpb.ServiceDescriptorProto
}

type BaseCode struct {
	Req          *pluginpb.CodeGeneratorRequest
	File         *descriptorpb.FileDescriptorProto
	Service      *descriptorpb.ServiceDescriptorProto
	Namespace    *Namespace
	Embed        embed.FS
	Template     *template.Template
	TemplateFile string
}

func NewBaseCode(req *pluginpb.CodeGeneratorRequest, file *descriptorpb.FileDescriptorProto, service *descriptorpb.ServiceDescriptorProto, template *template.Template, templateFile string) *BaseCode {
	return &BaseCode{
		Req:          req,
		File:         file,
		Service:      service,
		Namespace:    NewNamespace(PHP{}, req, file, service),
		Embed:        mt,
		Template:     template,
		TemplateFile: templateFile,
	}
}

func (b BaseCode) Body() (string, error) {
	out := bytes.NewBuffer(nil)
	data := value{
		Namespace: b.Namespace,
		File:      b.File,
		Service:   b.Service,
	}
	t, err := b.Template.ParseFS(b.Embed, b.TemplateFile)
	if err != nil {
		return "", err
	}
	err = t.Execute(out, data)
	return out.String(), nil
}

type ClientCode struct {
	*BaseCode
}

func NewClientCode(req *pluginpb.CodeGeneratorRequest, file *descriptorpb.FileDescriptorProto, service *descriptorpb.ServiceDescriptorProto, ns *Namespace) *ClientCode {
	tpl := template.New("client.tpl").Funcs(template.FuncMap{
		"client": func(name *string) string {
			return ns.p.Identifier(*name, "client")
		},
		"name": resolveNamespaceFunc(),
	})
	return &ClientCode{
		BaseCode: NewBaseCode(req, file, service, tpl, "template/client.tpl"),
	}
}

type InterfaceCode struct {
	*BaseCode
}

func NewInterfaceCode(req *pluginpb.CodeGeneratorRequest, file *descriptorpb.FileDescriptorProto, service *descriptorpb.ServiceDescriptorProto, ns *Namespace) *InterfaceCode {
	tpl := template.New("client.tpl").Funcs(template.FuncMap{
		"interface": func(name *string) string {
			return ns.p.Identifier(*name, "interface")
		},
		"name": resolveNamespaceFunc(),
	})
	return &InterfaceCode{
		BaseCode: NewBaseCode(req, file, service, tpl, "template/service_interface.tpl"),
	}
}

type ServiceCode struct {
	*BaseCode
}

func NewServiceCode(req *pluginpb.CodeGeneratorRequest, file *descriptorpb.FileDescriptorProto, service *descriptorpb.ServiceDescriptorProto, ns *Namespace) *ServiceCode {
	tpl := template.New("client.tpl").Funcs(template.FuncMap{
		"service": func(name *string) string {
			return ns.p.Identifier(*name, "service")
		},
		"interface": func(name *string) string {
			return ns.p.Identifier(*name, "interface")
		},
		"name": resolveNamespaceFunc(),
	})
	return &ServiceCode{
		BaseCode: NewBaseCode(req, file, service, tpl, "template/service.tpl"),
	}
}

func resolveNamespaceFunc() func(ns *Namespace, name *string) string {
	return func(ns *Namespace, name *string) string {
		return ns.resolve(name)
	}
}
