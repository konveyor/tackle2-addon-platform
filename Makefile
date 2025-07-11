GOPATH    ?= $(HOME)/go
GOBIN     ?= $(GOPATH)/bin
IMG       ?= quay.io/konveyor/tackle2-addon-platform:latest
CMD       ?= bin/addon
AddonDir  ?= /tmp/addon
GOIMPORTS = $(GOBIN)/goimports

PKG = ./cmd/...

PKGDIR = $(subst /...,,$(PKG))


cmd: fmt vet
	go build -ldflags="-w -s" -o ${CMD} github.com/konveyor/tackle2-addon-platform/cmd

image-docker:
	docker build -t ${IMG} .

image-podman:
	podman build -t ${IMG} .

run: cmd
	mkdir -p ${AddonDir}
	$(eval cmd := $(abspath ${CMD}))
	cd ${AddonDir};${cmd}

fmt: $(GOIMPORTS)
	$(GOIMPORTS) -w $(PKGDIR)

vet:
	go vet $(PKG)

test:
	go test -count=1 -v ./cmd/...

# Ensure goimports installed.
$(GOIMPORTS):
	go install golang.org/x/tools/cmd/goimports@v0.24
