.PHONY: all

GOCMD=go
GOBUILD=$(GOCMD) build -tags '!debug' -gcflags "all=-N -l" -ldflags "-s -w"
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean

DIST_DIR=./build/dist
PACKAGE=./build/panabit-ddns-go.tar.gz


all: clean build package

clean:
	rm -rf $(DIST_DIR)

build: build-ctl build-cgi

build-ctl:
	GOOS=linux GOARCH=arm64 $(GOBUILD) -o $(DIST_DIR)/appctl -v ./cmd
build-cgi:
	GOOS=linux GOARCH=arm64 $(GOBUILD) -o $(DIST_DIR)/web/cgi/webmain -v ./cmd/cgi

DDNSGO=./static/bin/ddns-go
DDNSGO_PATH=./static/bin
DDNSGO_TARBALL=./static/bin/ddns-go.tar.gz
DDNSGO_URL=https://github.com/jeessy2/ddns-go/releases/download/v5.6.6/ddns-go_5.6.6_linux_arm64.tar.gz

package: $(DDNSGO)
	cp -r ./static/* $(DIST_DIR)
	chmod +x $(DIST_DIR)/appctrl
	chmod +x $(DIST_DIR)/afterinstall
	tar -czvf $(PACKAGE) -C $(DIST_DIR) --exclude='.gitkeep' --exclude='LICENSE' --exclude='README.md' .

$(DDNSGO):
	wget -O $(DDNSGO_TARBALL) $(DDNSGO_URL)
	tar -xzvf $(DDNSGO_TARBALL) -C $(DDNSGO_PATH)
	rm $(DDNSGO_TARBALL)
