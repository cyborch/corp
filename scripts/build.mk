.PHONY: corp corp-clean

corp: $(ROOT_DIR)bin/corp

$(ROOT_DIR)bin/corp: $(wildcard cmd/corp/*.go)
	mkdir -p $(ROOT_DIR)
	go build -o $(ROOT_DIR)bin/corp $<

corp-clean:
	if [ $(ROOT_DIR) == "" ] ; then exit 1 ; fi
	if [ -f $(ROOT_DIR)bin/corp ] ; then rm $(ROOT_DIR)bin/corp ; fi
