SUB_BIN   = sub_bin
PWD       = $(shell pwd)
REPO      = $(shell basename $(PWD))
CMDPATH   = ./cmd/$(REPO)/
GOVERSION = $(shell go version)
GOOS      = $(word 1,$(subst /, ,$(lastword $(GOVERSION))))
GOARCH    = $(word 2,$(subst /, ,$(lastword $(GOVERSION))))
HAS_GLIDE = $(shell which glide)
XC_OS     = "darwin linux windows"
XC_ARCH   = "amd64"
VERSION   = $(patsubst "%",%,$(lastword $(shell grep 'version =' cmd/battery/main.go)))
RELEASE   = ./releases/$(VERSION)

GITHUB_USERNAME = "Code-Hex"

rm-build:
	@rm -rf build

rm-releases:
	@rm -rf releases

rm-all: rm-build rm-releases

release: all
	@mkdir -p $(RELEASE)
	@for DIR in $(shell ls ./build/$(VERSION)/) ; do \
		echo Processing in build/$(VERSION)/$$DIR; \
		cd $(PWD); \
		cp README.md ./build/$(VERSION)/$$DIR; \
		cp LICENSE ./build/$(VERSION)/$$DIR; \
		tar -cjf ./$(RELEASE)/$(REPO)_$(VERSION)_$$DIR.tar.bz2 -C ./build/$(VERSION) $$DIR; \
		tar -czf ./$(RELEASE)/$(REPO)_$(VERSION)_$$DIR.tar.gz -C ./build/$(VERSION) $$DIR; \
		cd build/$(VERSION); \
		zip -9 $(PWD)/$(RELEASE)/$(REPO)_$(VERSION)_$$DIR.zip $$DIR/*; \
	done

prepare-github: github-token
	@echo "'github-token' file is required"

release-upload: prepare-github release
	@echo "Uploading..."
	@ghr -u $(GITHUB_USERNAME) -t $(shell cat github-token) --draft --replace $(VERSION) $(RELEASE)

all:
	@PATH=$(SUB_BIN)/$(GOOS)/$(GOARCH):$(PATH)
	@gox -os=$(XC_OS) -arch=$(XC_ARCH) -output="build/$(VERSION)/{{.OS}}_{{.Arch}}/{{.Dir}}" $(CMDPATH)

