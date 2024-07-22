<?php

declare(strict_types=1);

# Generated by the protocol buffer compiler (for Phluxor). DO NOT EDIT!
# source: {{ .File.Name }}
{{ $ns := .Namespace -}}
{{if $ns.Namespace}}
namespace {{ $ns.Namespace }};
{{end}}
use Phluxor\GRPC;
{{- range $n := $ns.Import}}
use {{ $n }};
{{- end}}

{{ $package := .File.Package -}}
{{ $svc := .Service.Name -}}
{{ $once := 0 -}}

class {{ .Service.Name | client }} extends GRPC\BaseStub
{
{{- range $m := .Service.Method}}
{{if $m.ServerStreaming}}
	/**
    * @param {{ name $ns $m.InputType }} $request
    * @return {{ name $ns $m.OutputType }}
    *
    * @throws GRPC\Exception\InvokeException
    */
    public function {{ $m.Name }}({{ name $ns $m.InputType }} $request, $metadata = []): {{ name $ns $m.OutputType }} {
    	return $this->_serverStreamRequest('/{{ $package }}.{{ $svc }}/{{ $m.Name }}',
        $request,
        ['\{{ $ns.Namespace }}\{{ name $ns $m.OutputType }}', 'decode'],
        $metadata);
    }
{{if eq $once 0}}
	public function getNext(): object {
	    return $this->_getData();
	}
{{end -}}
{{ $once = 1}}
{{- else}}
    /**
    * @param {{ name $ns $m.InputType }} $request
    * @return {{ name $ns $m.OutputType }}
    *
    * @throws GRPC\Exception\InvokeException
    */
    public function {{ $m.Name }}({{ name $ns $m.InputType }} $request, $metadata = []): {{ name $ns $m.OutputType }} {
    	return $this->_simpleRequest('/{{ $package }}.{{ $svc }}/{{ $m.Name }}',
        $request,
        ['\{{ $ns.Namespace }}\{{ name $ns $m.OutputType }}', 'decode'],
        $metadata);
    }
{{end -}}

{{end -}}
}
