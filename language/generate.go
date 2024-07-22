package language

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

type PHPCode struct {
	PHP
}

func NewPHPCode() *PHPCode {
	return &PHPCode{PHP{}}
}

func (p *PHPCode) Generate(request *pluginpb.CodeGeneratorRequest) *pluginpb.CodeGeneratorResponse {
	response := &pluginpb.CodeGeneratorResponse{}
	for _, file := range request.ProtoFile {
		for _, service := range file.Service {
			response.File = append(response.File, p.generateInterface(request, file, service))
			response.File = append(response.File, p.generateService(request, file, service))
			response.File = append(response.File, p.generateClient(request, file, service))
		}
	}
	return response
}

func (p *PHPCode) generateClient(
	req *pluginpb.CodeGeneratorRequest,
	file *descriptorpb.FileDescriptorProto,
	service *descriptorpb.ServiceDescriptorProto) *pluginpb.CodeGeneratorResponse_File {

	nic := NewClientCode(req, file, service, NewNamespace(p.PHP, req, file, service))
	code, _ := nic.Body()
	i := &Client{p: p.PHP}
	return &pluginpb.CodeGeneratorResponse_File{
		Name:    proto.String(i.Filename(file, service.Name)),
		Content: proto.String(code),
	}
}

func (p *PHPCode) generateInterface(
	req *pluginpb.CodeGeneratorRequest,
	file *descriptorpb.FileDescriptorProto,
	service *descriptorpb.ServiceDescriptorProto) *pluginpb.CodeGeneratorResponse_File {

	nic := NewInterfaceCode(req, file, service, NewNamespace(p.PHP, req, file, service))
	code, _ := nic.Body()
	i := &InterfaceName{p: p.PHP}
	return &pluginpb.CodeGeneratorResponse_File{
		Name:    proto.String(i.Filename(file, service.Name)),
		Content: proto.String(code),
	}
}

func (p *PHPCode) generateService(
	req *pluginpb.CodeGeneratorRequest,
	file *descriptorpb.FileDescriptorProto,
	service *descriptorpb.ServiceDescriptorProto) *pluginpb.CodeGeneratorResponse_File {

	nic := NewServiceCode(req, file, service, NewNamespace(p.PHP, req, file, service))
	code, _ := nic.Body()
	i := &ServiceName{p: p.PHP}
	return &pluginpb.CodeGeneratorResponse_File{
		Name:    proto.String(i.Filename(file, service.Name)),
		Content: proto.String(code),
	}
}
