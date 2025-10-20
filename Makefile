include make.conf

EXECUTABLE=go-mpbt
TEMPFILES=*.tmp .*.tmp .tmp swagger.json blockstore/.tmp
# SUBDIRS := tests core/wire cmd/perseus tests/rsatest

#test: compile
#	@if $(GO) test -v $(PACKAGE)/... ; then echo "=== Test okay ===" ; else echo " ==== self-test failed === "; exit 1 ; fi

compile:
	$(MAKE) -C cmd
	for d in $(SUBDIRS) ; do $(MAKE) -C $$d compile ; done

fmt:
	$(GO) fmt $(PACKAGE)/...

clean:
	for d in $(SUBDIRS) ; do $(MAKE) -C $$d clean ; done
	rm -Rf $(EXECUTABLE) $(TEMPFILES)

swagger:
	swagger generate spec -o ./swagger.json

swagger-serve: swagger
	swagger serve -F=swagger swagger.json

.PHONY: swagger swagger-serve proto

run:
	$(MAKE) -C cmd
	./cmd/mpbt-builder/mpbt-builder -root . -solution cf/xlibre/solutions/devuan.yaml
