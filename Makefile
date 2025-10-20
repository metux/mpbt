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
	./cmd/mpbt-builder/mpbt-builder -root . -solution cf/xlibre/solutions/devuan.yaml

.PHONY: compile fmt clean run
