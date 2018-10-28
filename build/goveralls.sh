#!/bin/bash
set -e
set -x
golint -set_exit_status .
go vet .
echo "mode: count" > c1.out
go test -coverprofile=c.out .
go test -covermode=count -coverprofile=c1.out .
goveralls -coverprofile=c1.out -service=travis-ci -repotoken $COVERALLS_TOKEN

