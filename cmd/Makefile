include ../make.conf

SUBDIRS := mpbt-builder

compile:
	for i in $(SUBDIRS) ; do $(MAKE) -C $$i ; done

clean:
	for i in $(SUBDIRS) ; do $(MAKE) -C $$i clean ; done
	rm -f $(EXECUTABLE)
