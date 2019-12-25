GO=go
NAME := tjdoe
VERSION := 1.0.0

all: test build

deps:
	$(GO) get github.com/mattn/goveralls

setup: update_version
	git submodule update --init

update_version:
	@for i in README.md docs/content/_index.md; do\
	    sed -e 's!Version-[0-9.]*-yellowgreen!Version-${VERSION}-yellowgreen!g' -e 's!tag/v[0-9.]*!tag/v${VERSION}!g' $$i > a ; mv a $$i; \
	done
	@sed 's/const VERSION = .*/const VERSION = "${VERSION}"/g' cmd/tjdoe/main.go > a
	@mv a cmd/tjdoe/main.go
	@echo "Replace version to \"${VERSION}\""

test: setup
	$(GO) test -covermode=count -coverprofile=coverage.out $$(go list ./... | grep -v vendor)

build: setup
	$(GO) build -o $(NAME) -v cmd/tjdoe/main.go

clean:
	$(GO) clean
	rm -rf $(NAME)
