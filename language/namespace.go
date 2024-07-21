package language

import (
	"bytes"
	"fmt"
	"strings"

	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

// Namespace defines namespace struct.
type Namespace struct {
	p PHP
	// Package defines file package.
	Package string
	// Namespace defines namespace.
	Namespace string
	// Import declares what namespaces to be imported
	Import map[string]string
}

func NewNamespace(
	p PHP,
	req *pluginpb.CodeGeneratorRequest,
	file *descriptorpb.FileDescriptorProto,
	service *descriptorpb.ServiceDescriptorProto) *Namespace {
	ns := &Namespace{
		Package:   *file.Package,
		Namespace: p.Namespace(file.Package, "\\"),
		Import:    make(map[string]string),
	}
	if file.Options != nil && file.Options.PhpNamespace != nil {
		ns.Namespace = *file.Options.PhpNamespace
	}
	for k := range service.Method {
		ns.importMessage(req, service.Method[k].InputType)
		ns.importMessage(req, service.Method[k].OutputType)
	}
	return ns
}

// importMessage registers new import message namespace (only the namespace).
func (ns *Namespace) importMessage(req *pluginpb.CodeGeneratorRequest, msg *string) {
	if msg == nil {
		return
	}
	chunks := strings.Split(*msg, ".")
	pkg := strings.Join(chunks[:len(chunks)-1], ".")
	result := bytes.NewBuffer(nil)
	for _, p := range chunks[:len(chunks)-1] {
		result.WriteString(ns.p.Identifier(p, ""))
		result.WriteString(`\`)
	}
	if pkg == "."+ns.Package {
		// root package
		return
	}
	for _, f := range req.ProtoFile {
		if pkg == "."+*f.Package {
			if f.Options != nil && f.Options.PhpNamespace != nil {
				// custom imported namespace
				ns.Import[pkg] = *f.Options.PhpNamespace
				return
			}
		}
	}

	ns.Import[pkg] = strings.Trim(result.String(), `\`)
}

// resolve message alias
func (ns *Namespace) resolve(msg *string) string {
	chunks := strings.Split(*msg, ".")
	pkg := strings.Join(chunks[:len(chunks)-1], ".")
	if pkg == "."+ns.Package {
		return ns.p.Identifier(chunks[len(chunks)-1], "")
	}
	for iPkg, s := range ns.Import {
		if pkg == iPkg {
			nsChunks := strings.Split(s, `\`)
			identifier := ns.p.Identifier(chunks[len(chunks)-1], "")
			return fmt.Sprintf(
				`%s\%s`,
				nsChunks[len(nsChunks)-1],
				ns.p.resolveReserved(identifier, pkg),
			)
		}
	}
	return "\\" + ns.p.Namespace(msg, "\\")
}
