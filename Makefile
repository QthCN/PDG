.PHONY: all be clean devstack

GOPATH :=
ifeq ($(OS),Windows_NT)
	GOPATH := $(CURDIR)
else
	GOPATH := $(CURDIR)
endif

export GOPATH
export CGO_ENABLED := 0
#export GOOS := linux
export GOARCH := amd64


all: be

be:
	go install be/cmd/pdg

clean:
	rm -rf bin
	rm -rf pkg