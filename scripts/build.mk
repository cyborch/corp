.PHONY: corp corp-clean

GO_SOURCES := $(shell find $(ROOT_DIR)cmd $(ROOT_DIR)internal -name '*.go')

corp: $(ROOT_DIR)bin/corp

$(ROOT_DIR)bin/corp: $(GO_SOURCES)
	mkdir -p $(ROOT_DIR)bin
	go build -o $(ROOT_DIR)bin/corp $(wildcard cmd/corp/*.go)

corp-clean:
	if [ $(ROOT_DIR) == "" ] ; then exit 1 ; fi
	if [ -f $(ROOT_DIR)bin/corp ] ; then rm $(ROOT_DIR)bin/corp ; fi
