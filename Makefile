# VARIABLES
PACKAGES := $(shell go list ./...)
PROJECT := $(shell basename ${PWD})
LOCAL_BIN := $(CURDIR)/bin
USER := rptl8sr
EMAIL := $(USER)@gmail.com

# GITHUB
.PHONY: git-init
git-init:
	gh repo create $(PROJECT) --private
	git init
	git config user.name "$(USER)"
	git config user.email "$(EMAIL)"
	git add Makefile go.mod
	git commit -m "Init commit"
	git remote add origin git@github.com:$(USER)/$(PROJECT).git
	git remote -v
	git push -u origin master


BN ?= dev
# make git-checkout BN=dev
.PHONY: git-checkout
git-checkout:
	git checkout -b $(BN)


# BUILDING
.PHONY: build
build: build-templ test build-app

.PHONY: build-templ
build-templ:
	templ generate


.PHONY: test
test:
	go test -race -cover $(PACKAGES)


.PHONY: build-app
build-app:
	go build -o ./tmp/main -v .


# WATCHING
.PHONY: watch
watch:
	$(MAKE) -j2 watch-templ watch-app



.PHONY: watch-templ
watch-templ:
	templ generate \
	--watch \
	--proxy="http://localhost:8080" \
	--open-browser=false


.PHONY: watch-app
watch-app:
	air