# SPDX-License-Identifier: AGPL-3.0-or-later
include ../../make.conf

EXECUTABLE := mpbt-builder

compile:
	GOOS=linux GOARCH=386 CGO_ENABLED=0 $(GO) build $(GOTAGS) -trimpath -ldflags="-s -w" .

clean:
	rm -f $(EXECUTABLE)
