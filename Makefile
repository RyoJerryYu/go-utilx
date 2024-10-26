.PHONY: tidy
tidy:
	@find . -name go.mod -print | while read -r gomod_path; do \
		dir_path=$$(dirname "$$gomod_path"); \
		echo "Executing 'go mod tidy' in directory: $$dir_path"; \
		(cd "$$dir_path"  && GOPROXY=$(GOPROXY) go get -u ./... && GOPROXY=$(GOPROXY) go mod tidy) || exit 1; \
	done

.PHONY: test
test:
	@go test -v ./...