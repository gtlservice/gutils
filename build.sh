#!/bin/bash

go build github.com/gtlservice/gutils/algorithm
go build github.com/gtlservice/gutils/container
go build github.com/gtlservice/gutils/logger
go build github.com/gtlservice/gutils/network
go build github.com/gtlservice/gutils/system
go build github.com/gtlservice/gutils/compress/tarlib
go build github.com/gtlservice/gutils/http
go build github.com/gtlservice/gutils/flocker
go build

go install github.com/gtlservice/gutils/algorithm
go install github.com/gtlservice/gutils/container
go install github.com/gtlservice/gutils/logger
go install github.com/gtlservice/gutils/network
go install github.com/gtlservice/gutils/system
go install github.com/gtlservice/gutils/compress/tarlib
go install github.com/gtlservice/gutils/http
go install github.com/gtlservice/gutils/flocker
go install