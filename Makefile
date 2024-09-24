PLATFORMS := $(shell go tool dist list | grep -E 'linux|windows|darwin' | grep -E 'amd64|arm64|386')
SOURCE := proxy.go
DEST := $(SOURCE:.go=)
LDFLAGS := -s -w
USE_UPX := 1

DELIM = $(subst /, ,$@)
OS = $(word 1, $(DELIM))
ARCH = $(word 2, $(DELIM))

release: $(PLATFORMS)

$(PLATFORMS): $(SOURCE)
	GOOS=$(OS) GOARCH=$(ARCH) go build -ldflags "$(LDFLAGS)" -o $(DEST)-$(OS)-$(ARCH) $(SOURCE) 
	if [ $(USE_UPX) -eq 1 ] && [ "$(OS)" != "darwin" ]; then upx $(DEST)-$(OS)-$(ARCH); fi

build: $(SOURCE)
	go build -ldflags "$(LDFLAGS)" -o $(DEST) $(SOURCE)
	if [ $(USE_UPX) -eq 1 ]; then upx $(DEST); fi

.PHONY: release build $(PLATFORMS)
