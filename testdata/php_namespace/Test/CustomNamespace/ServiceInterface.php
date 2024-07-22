<?php

declare(strict_types=1);

# Generated by the protocol buffer compiler (for Phluxor). DO NOT EDIT!
# source: php_namespace/service.proto

namespace Test\CustomNamespace;

use Phluxor\GRPC;

interface ServiceInterface extends GRPC\ServiceInterface
{
    // GRPC specific service name.
    public const NAME = "testPhpNamespace.Service";

    /**
    * @param GRPC\ContextInterface $ctx
    * @param SimpleMessage $in
    * @return SimpleMessage
    *
    * @throws GRPC\Exception\InvokeException
    */
    public function SimpleMethod(GRPC\ContextInterface $ctx, SimpleMessage $in): SimpleMessage;
}
