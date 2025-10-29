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
            -workdir WORK.tmp \
            -solution cf/xlibre/solutions/devuan.yaml \
            -project-define xlibre_git=git@github.com:X11Libre \
            build

.PHONY: compile fmt clean run
