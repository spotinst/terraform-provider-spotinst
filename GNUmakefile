TEST?=./...
PKGNAME?=spotinst
VERSION?=$(shell grep -oP '(?<=Version = ).+' version/version.go | xargs)
RELEASE?=v$(VERSION)
SUCCESSFUL_TESTS_RUN?=$(shell bash -c 'read -s -p "last successful tests ran on jenkins build number: " pwd; echo $$pwd')

default: build

.PHONY: build
build: fmtcheck
	go install

.PHONY: test
test: fmtcheck
	go test $(TEST) -timeout=30s -parallel=4

.PHONY: testacc
testacc: fmtcheck
	TF_ACC=1 go test $(TEST) -v -count 1 -parallel 20 $(TESTARGS) -timeout 120m

.PHONY: testcompile
testcompile:
	@if [ "$(TEST)" = "./..." ]; then \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./$(PKGNAME)"; \
		exit 1; \
	fi
	go test -c $(TEST) $(TESTARGS)

.PHONY: vet
vet:
	go vet ./...

.PHONY: fmt
fmt:
	@gofmt -s -w ./$(PKGNAME)

.PHONY: fmtcheck
fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

.PHONY: depscheck
depscheck:
	@echo "==> Checking source code with go mod tidy..."
	@go mod tidy
	@git diff --exit-code -- go.mod go.sum || \
		(echo; echo "Unexpected difference in go.mod/go.sum files. Run 'go mod tidy' command or revert any go.mod/go.sum changes and commit."; exit 1)

.PHONY: docs
docs: tools
	@sh -c "'$(CURDIR)/scripts/generate-docs.sh'"

.PHONY: docscheck
docscheck: docs
	@tfplugindocs validate

.PHONY: lint
lint:
	@echo "==> Checking source code against linters..."
	@golint ./$(PKGNAME)/...
	@tfproviderlint \
		-c 1 \
		-AT001 \
		-AT002 \
		-S001 \
		-S002 \
		-S003 \
		-S004 \
		-S005 \
		-S007 \
		-S008 \
		-S009 \
		-S010 \
		-S011 \
		-S012 \
		-S013 \
		-S014 \
		-S015 \
		-S016 \
		-S017 \
		-S019 \
		./$(PKGNAME)

.PHONY: tools
tools:
	@go generate -tags tools tools.go

.PHONY: release
release:
	@git commit -a -m "chore(release): $(RELEASE)" -m "last successful tests ran on jenkins build number: $(SUCCESSFUL_TESTS_RUN)"
	@git tag -f -m    "chore(release): $(RELEASE)" $(RELEASE)
	@git push --follow-tags
