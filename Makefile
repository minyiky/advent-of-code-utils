VERSION=0.3.0

version:
	@echo "v${VERSION}"

tag:
	git tag "v${VERSION}"

test-unit:
	go test --coverprofile=cover.out --count=1 ./pkg/...

ci-test-unit:
	@go test --json --coverprofile=cover.out --count=1 ./pkg/...
