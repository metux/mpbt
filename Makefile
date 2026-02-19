include make.conf

compile:
	$(MAKE) -C cmd
	for d in $(SUBDIRS) ; do $(MAKE) -C $$d compile ; done

fmt:
	$(GO) fmt $(PACKAGE)/...

clean:
	for d in $(SUBDIRS) ; do $(MAKE) -C $$d clean ; done

run:
	$(MAKE) -C cmd
	./cmd/mpbt-builder/mpbt-builder \
            -root . \
            -workdir WORK \
            -solution cf/xlibre/solutions/devuan.yaml \
            -project-define xlibre_git=git@github.com:X11Libre \
            build

run2:
	$(GO) run scripts/test-build-xlibre.go

.PHONY: compile fmt clean run

GOPATH := $(shell go env GOPATH)

install: compile
	mkdir -p $(GOPATH)/bin
	cp cmd/mpbt-builder/mpbt-builder $(GOPATH)/bin
