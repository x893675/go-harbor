.PHONY: fmt
fmt:
	gofmt -s -w ./


.PHONY: checkfmt
checkfmt:
	@echo checking gofmt...
	@res=$$(gofmt -d -e -s $$(find . -type d \( -path ./src/vendor -o -path ./tests \) -prune -o -name '*.go' -print)); \
	if [ -n "$${res}" ]; then \
		echo checking gofmt fail... ; \
		echo "$${res}"; \
		exit 1; \
	fi


.PHONY: test
ifeq "$(strip $(shell go env GOARCH))" "amd64"
RACE_FLAG := -race
endif
test:
	go test $(RACE_FLAG) -vet all $(go list ./... | grep -v /example/)
