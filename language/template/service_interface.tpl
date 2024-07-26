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

interface {{ .Service.Name | interface }} extends GRPC\ServiceInterface
{
    public const string NAME = "/{{ .File.Package }}.{{ .Service.Name }}";{{ "\n" }}
{{- range $m := .Service.Method}}
    /**
     * @param GRPC\ContextInterface $ctx
     * @param {{ name $ns $m.InputType }} $request
     * @return {{ name $ns $m.OutputType }}
     *
     * @throws GRPC\Exception\InvokeException
     */
    public function {{ $m.Name }}(GRPC\ContextInterface $ctx, {{ name $ns $m.InputType }} $request): {{ name $ns $m.OutputType }}; // @phpcs:ignore
{{end -}}
}
