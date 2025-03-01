package language

import (
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

const expectedClientCode = `<?php

declare(strict_types=1);

# Generated by the protocol buffer compiler (for Phluxor). DO NOT EDIT!
# source: testdata/service_filename.proto

namespace Test\CustomNamespace;

use Google\Protobuf\Internal\Message;
use Phluxor\GRPC;
use Test;

class TestClient extends GRPC\BaseStub
{
    /**
     * @param Test\Request $request
     * @param array<string|int, mixed> $metadata
     * @return Test\Response
     *
     * @throws GRPC\Exception\InvokeException|\Exception
     */
    public function Test(Test\Request $request, array $metadata = []): Test\Response // @phpcs:ignore
    {
    	return $this->serverStreamRequest('/test.Test/Test',
        $request,
        ['\Test\CustomNamespace\Test\Response', 'decode'],
        $metadata);
    }

    /**
     * @throws \Exception
     */
    public function getNext(): ?Message
    {
        return $this->getData();
    }
}
`

func TestClientCode_Body(t *testing.T) {
	fdp := &descriptorpb.FileDescriptorProto{
		Name: proto.String("testdata/service_filename.proto"),
		Options: &descriptorpb.FileOptions{
			PhpNamespace: proto.String("Test\\CustomNamespace"),
		},
		Package: proto.String("test"),
	}
	req := &pluginpb.CodeGeneratorRequest{
		ProtoFile: []*descriptorpb.FileDescriptorProto{fdp},
	}
	sdp := &descriptorpb.ServiceDescriptorProto{
		Name: proto.String("Test"),
		Method: []*descriptorpb.MethodDescriptorProto{
			{
				Name:            proto.String("Test"),
				InputType:       proto.String("Test.Request"),
				OutputType:      proto.String("Test.Response"),
				ServerStreaming: proto.Bool(false),
			},
		},
	}
	cc := NewClientCode(req, fdp, sdp, NewNamespace(PHP{}, req, fdp, sdp))
	code, err := cc.Body()
	if err != nil {
		t.Error(err)
	}
	if code != expectedClientCode {
		t.Errorf("ClientCode.Body() failed: %s", code)
	}
}

const expectedInterfaceCode = `<?php

declare(strict_types=1);

# Generated by the protocol buffer compiler (for Phluxor). DO NOT EDIT!
# source: testdata/service_filename.proto

namespace Test\CustomNamespace;

use Phluxor\GRPC;
use Test;

interface TestInterface extends GRPC\ServiceInterface
{
    public const string NAME = "/test.Test";

    /**
     * @param GRPC\ContextInterface $ctx
     * @param Test\Request $request
     * @return Test\Response
     *
     * @throws GRPC\Exception\InvokeException
     */
    public function Test(GRPC\ContextInterface $ctx, Test\Request $request): Test\Response; // @phpcs:ignore
}
`

func TestNewInterfaceCode(t *testing.T) {
	fdp := &descriptorpb.FileDescriptorProto{
		Name: proto.String("testdata/service_filename.proto"),
		Options: &descriptorpb.FileOptions{
			PhpNamespace: proto.String("Test\\CustomNamespace"),
		},
		Package: proto.String("test"),
	}
	req := &pluginpb.CodeGeneratorRequest{
		ProtoFile: []*descriptorpb.FileDescriptorProto{fdp},
	}
	sdp := &descriptorpb.ServiceDescriptorProto{
		Name: proto.String("Test"),
		Method: []*descriptorpb.MethodDescriptorProto{
			{
				Name:            proto.String("Test"),
				InputType:       proto.String("Test.Request"),
				OutputType:      proto.String("Test.Response"),
				ServerStreaming: proto.Bool(false),
			},
		},
	}
	cc := NewInterfaceCode(req, fdp, sdp, NewNamespace(PHP{}, req, fdp, sdp))
	code, err := cc.Body()
	if err != nil {
		t.Error(err)
	}
	if code != expectedInterfaceCode {
		t.Errorf("InterfaceCode.Body() failed: %s", code)
	}
}

const expectedServiceCode = `<?php

declare(strict_types=1);

# Generated by the protocol buffer compiler (for Phluxor). DO NOT EDIT!
# source: testdata/service_filename.proto

namespace Test\CustomNamespace;

use Phluxor\GRPC;
use Test;

class TestService implements TestInterface
{
    /**
     * @param GRPC\ContextInterface $ctx
     * @param Test\Request $request
     * @return Test\Response
     *
     * @throws GRPC\Exception\InvokeException
     */
    public function Test(GRPC\ContextInterface $ctx, Test\Request $request): Test\Response // @phpcs:ignore
    {
    	// your code
    }
}
`

func TestNewServiceCode(t *testing.T) {
	fdp := &descriptorpb.FileDescriptorProto{
		Name: proto.String("testdata/service_filename.proto"),
		Options: &descriptorpb.FileOptions{
			PhpNamespace: proto.String("Test\\CustomNamespace"),
		},
		Package: proto.String("test"),
	}
	req := &pluginpb.CodeGeneratorRequest{
		ProtoFile: []*descriptorpb.FileDescriptorProto{fdp},
	}
	sdp := &descriptorpb.ServiceDescriptorProto{
		Name: proto.String("Test"),
		Method: []*descriptorpb.MethodDescriptorProto{
			{
				Name:            proto.String("Test"),
				InputType:       proto.String("Test.Request"),
				OutputType:      proto.String("Test.Response"),
				ServerStreaming: proto.Bool(false),
			},
		},
	}
	cc := NewServiceCode(req, fdp, sdp, NewNamespace(PHP{}, req, fdp, sdp))
	code, err := cc.Body()
	if err != nil {
		t.Error(err)
	}
	if code != expectedServiceCode {
		t.Errorf("ServiceCode.Body() failed: %s", code)
	}
}
