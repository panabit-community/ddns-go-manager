.PHONY: all clean build package

GOCMD=go
GOBUILD=$(GOCMD) build -tags '!debug' -gcflags "all=-N -l" -ldflags "-s -w"
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean

DIST_DIR=./build/dist
PACKAGE_AMD64=./build/panabit-ddns-go-amd64.tar.gz
PACKAGE_ARM64=./build/panabit-ddns-go-arm64.tar.gz

all: clean package

clean:
	rm -rf $(DIST_DIR)

build:
	$(MAKE) build-ctl ARCH=$(ARCH)
	$(MAKE) build-cgi ARCH=$(ARCH)

build-ctl:
	GOOS=linux GOARCH=$(ARCH) $(GOBUILD) -o $(DIST_DIR)/appctl -v ./cmd
build-cgi:
	GOOS=linux GOARCH=$(ARCH) $(GOBUILD) -o $(DIST_DIR)/web/cgi/webmain -v ./cmd/cgi

CA_BUNDLE=./static/certs/ca-bundle.crt

DDNSGO=./static/bin/ddns-go
DDNSGO_PATH=./static/bin
DDNSGO_TARBALL=./static/bin/ddns-go.tar.gz
DDNSGO_URL_AMD64=https://github.com/jeessy2/ddns-go/releases/download/v5.7.1/ddns-go_5.7.1_linux_x86_64.tar.gz
DDNSGO_URL_ARM64=https://github.com/jeessy2/ddns-go/releases/download/v5.7.1/ddns-go_5.7.1_linux_arm64.tar.gz

package: package-amd64 package-arm64

package-amd64: $(CA_BUNDLE)
	$(MAKE) build ARCH=amd64
	$(MAKE) package-multiarch DDNSGO_URL=$(DDNSGO_URL_AMD64) PACKAGE=$(PACKAGE_AMD64)

package-arm64: $(CA_BUNDLE)
	$(MAKE) build ARCH=arm64
	$(MAKE) package-multiarch DDNSGO_URL=$(DDNSGO_URL_ARM64) PACKAGE=$(PACKAGE_ARM64)

package-multiarch: $(CA_BUNDLE)
	curl -fsSLo $(DDNSGO_TARBALL) $(DDNSGO_URL)
	tar -xzvf $(DDNSGO_TARBALL) -C $(DDNSGO_PATH)
	rm $(DDNSGO_TARBALL)
	cp -r ./static/* $(DIST_DIR)
	chmod +x $(DIST_DIR)/appctrl
	chmod +x $(DIST_DIR)/afterinstall
	tar -czvf $(PACKAGE) -C $(DIST_DIR) --exclude='.gitkeep' --exclude='LICENSE' --exclude='README.md' .

$(CA_BUNDLE):
	curl -fsSLo $(CA_BUNDLE) https://curl.se/ca/cacert.pem
