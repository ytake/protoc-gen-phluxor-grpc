package language

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/types/descriptorpb"
)

type FilenameGenerator interface {
	Filename(file *descriptorpb.FileDescriptorProto, name *string) string
}

func (p PHP) DetectNamespace(file *descriptorpb.FileDescriptorProto) string {
	ns := p.Namespace(file.Package, "/")
	if file.Options != nil && file.Options.PhpNamespace != nil {
		ns = strings.ReplaceAll(*file.Options.PhpNamespace, `\`, `/`)
	}
	return ns
}

type Client struct {
	p PHP
}

func (f Client) Filename(file *descriptorpb.FileDescriptorProto, name *string) string {
	ns := f.p.DetectNamespace(file)
	return fmt.Sprintf("%s/%s.php", ns, f.p.Identifier(*name, "client"))
}

type InterfaceName struct {
	p PHP
}

func (f InterfaceName) Filename(file *descriptorpb.FileDescriptorProto, name *string) string {
	ns := f.p.DetectNamespace(file)
	return fmt.Sprintf("%s/%s.php", ns, f.p.Identifier(*name, "interface"))
}

type ServiceName struct {
	p PHP
}

func (f ServiceName) Filename(file *descriptorpb.FileDescriptorProto, name *string) string {
	ns := f.p.DetectNamespace(file)
	return fmt.Sprintf("%s/%s.php", ns, f.p.Identifier(*name, "service"))
}
