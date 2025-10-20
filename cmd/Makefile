include ../make.conf

EXECUTABLE := mpbt

compile:
	$(GO) build $(GOTAGS) .

clean:
	rm -f $(EXECUTABLE)

run: compile
	./cmd -root .. -solution ../cf/xlibre/solutions/devuan.yaml
