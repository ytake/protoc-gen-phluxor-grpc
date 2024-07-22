package language

import (
	"fmt"

	"google.golang.org/protobuf/types/descriptorpb"
)

type FilenameGenerator interface {
	Filename(file *descriptorpb.FileDescriptorProto, name *string) string
}

type Client struct {
	p PHP
}

func (f *Client) Filename(file *descriptorpb.FileDescriptorProto, name *string) string {
	ns := f.p.DetectNamespace(file)
	return fmt.Sprintf("%s/%s.php", ns, f.p.Identifier(*name, "client"))
}

type InterfaceName struct {
	p PHP
}

func (f *InterfaceName) Filename(file *descriptorpb.FileDescriptorProto, name *string) string {
	ns := f.p.DetectNamespace(file)
	return fmt.Sprintf("%s/%s.php", ns, f.p.Identifier(*name, "interface"))
}

type ServiceName struct {
	p PHP
}

func (f *ServiceName) Filename(file *descriptorpb.FileDescriptorProto, name *string) string {
	ns := f.p.DetectNamespace(file)
	return fmt.Sprintf("%s/%s.php", ns, f.p.Identifier(*name, "service"))
}
