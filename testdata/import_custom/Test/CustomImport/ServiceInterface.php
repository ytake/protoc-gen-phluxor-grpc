<?php

declare(strict_types=1);

# Generated by the protocol buffer compiler (for Phluxor). DO NOT EDIT!
# source: import_custom/service.proto

namespace Test\CustomImport;

use Phluxor\GRPC;
use Test\CustomImport\Message;

interface ServiceInterface extends GRPC\ServiceInterface
{
    // GRPC specific service name.
    public const NAME = "import.Service";

    /**
    * @param GRPC\ContextInterface $ctx
    * @param Message $in
    * @return Message
    *
    * @throws GRPC\Exception\InvokeException
    */
    public function SimpleMethod(GRPC\ContextInterface $ctx, Message $in): Message;

    /**
    * @param GRPC\ContextInterface $ctx
    * @param Message\Message $in
    * @return Message\Message
    *
    * @throws GRPC\Exception\InvokeException
    */
    public function ImportMethod(GRPC\ContextInterface $ctx, Message\Message $in): Message\Message;
}
