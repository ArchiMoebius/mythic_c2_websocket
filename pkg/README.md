# `pkg/`

This directory contains the core code which is invoked by the `/cmd/*/main.go` code.

## `pkg/api/`

This directory holds all of the gin / web api code

## `pkg/constants/`

There shouldn't be anything here but global constants...

## `pkg/config/`

Mailpipe and Mailpipe API config parsing code (most important is currently the transports code)

## `pkg/logger/`

Contains code to do console logging - but a file logger could be added...

## `pkg/mailpipe/`

This directory holds all of the mailpipe code (the part which accepts, parses, and then stores incoming emails)

## `pkg/storage/`

Currently, there is only redis storage of parsed emails - this could change...