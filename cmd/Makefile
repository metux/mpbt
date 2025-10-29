# SPDX-License-Identifier: AGPL-3.0-or-later
include ../make.conf

SUBDIRS := mpbt-builder

compile:
	for i in $(SUBDIRS) ; do $(MAKE) -C $$i ; done

clean:
	for i in $(SUBDIRS) ; do $(MAKE) -C $$i clean ; done
	rm -f $(EXECUTABLE)
