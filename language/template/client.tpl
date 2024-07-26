<?php

declare(strict_types=1);

# Generated by the protocol buffer compiler (for Phluxor). DO NOT EDIT!
# source: {{ .File.Name }}
{{ $ns := .Namespace -}}
{{if $ns.Namespace}}
namespace {{ $ns.Namespace }};
{{end}}
use Google\Protobuf\Internal\Message;
use Phluxor\GRPC;
{{- range $n := $ns.Import}}
use {{ $n }};
{{- end}}

{{ $package := .File.Package -}}
{{ $svc := .Service.Name -}}
{{ $once := 0 -}}

class {{ .Service.Name | client }} extends GRPC\BaseStub
{
{{- range $m := .Service.Method}}{{if $m.ServerStreaming}}
    /**
     * @param {{ name $ns $m.InputType }} $request
     * @param array<string|int, mixed> $metadata
     * @return {{ name $ns $m.OutputType }}
     *
     * @throws GRPC\Exception\InvokeException|\Exception
     */
    public function {{ $m.Name }}({{ name $ns $m.InputType }} $request, array $metadata = []): {{ name $ns $m.OutputType }} // @phpcs:ignore
    {
    	return $this->serverStreamRequest('/{{ $package }}.{{ $svc }}/{{ $m.Name }}',
        $request,
        ['\{{ $ns.Namespace }}\{{ name $ns $m.OutputType }}', 'decode'],
        $metadata);
    }
{{if eq $once 0}}
    /**
     * @throws \Exception
     */
    public function getNext(): ?Message
    {
        return $this->getData();
    }
{{end -}}
{{ $once = 1}}
{{- else}}
    /**
     * @param {{ name $ns $m.InputType }} $request
     * @param array<string|int, mixed> $metadata
     * @return {{ name $ns $m.OutputType }}
     *
     * @throws GRPC\Exception\InvokeException|\Exception
     */
    public function {{ $m.Name }}({{ name $ns $m.InputType }} $request, array $metadata = []): {{ name $ns $m.OutputType }} // @phpcs:ignore
    {
    	return $this->simpleRequest('/{{ $package }}.{{ $svc }}/{{ $m.Name }}',
        $request,
        ['\{{ $ns.Namespace }}\{{ name $ns $m.OutputType }}', 'decode'],
        $metadata);
    }
{{end -}}
{{end -}}
}
