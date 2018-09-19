ifneq (0, $(shell git tag -l --contains HEAD | wc -l | xargs))
	version = $(shell git describe --abbrev=0)
	commit = $(shell git rev-parse --short HEAD)
	ifeq ($(commit), $(shell git rev-list -n 1 $(version) | cut -c1-7))
		fullversion = $(version)
	else
		fullversion = $(version)-$(commit)
	endif
else
	fullversion = "v0.1.0-dev"
endif

stdbindir = $(GOPATH)/bin
ldflagsversion = -X main.rawVersion=$(fullversion)

install-std:
	go build -ldflags "$(ldflagsversion) -s -w" -o $(stdbindir)/iexplorer .

install:
	go build -race -tags "debug dev" -ldflags "$(ldflagsversion)" -o $(stdbindir)/iexplorer .
