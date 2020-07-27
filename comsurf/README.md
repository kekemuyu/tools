# Gofile

[![GoDoc](https://godoc.org/github.com/siadat/gofile/http?status.svg)](https://godoc.org/github.com/siadat/gofile/http)
[![Build Status](https://travis-ci.org/siadat/gofile.svg?branch=master)](https://travis-ci.org/siadat/gofile)

A non-blocking directory listing and file server.
It implementats HTTP/1.1 keepalive, chunked transfer, and byte range.

The HTTP server implementation provides a channel for writing chunked response. It could be used as a library. [Read the API](https://godoc.org/github.com/siadat/gofile/http).

![gofile](/../screenshots/screenshot-0.1.0.png?raw=true "gofile")

### Usage

    Usage: gofile port [dir]

Examples:

    gofile 8080
    gofile 8080 ~/public

### Install

    go get -u github.com/siadat/gofile

### HTTP/1.1 implementation checklist

- [x] GET and HEAD methods
- [x] Support keep-alive connections
- [x] Support chunked transfer encoding
- [x] Requests must include a `Host` header
- [x] Requests with `Connection: close` should be closed
- [x] Support for requests with absolute URLs
- [x] If-Modified-Since support
- [x] Byte range support
- [ ] Transparent response compression

### Hacking

Submit an issue or send a pull request.
Make sure you `./run-tests.bash` to test your patch.

### Thanks

Thanks @valyala for his feature suggestions. Thanks @maruel for reviewing the http package.
