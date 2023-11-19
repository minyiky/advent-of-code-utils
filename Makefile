VERSION=0.2.0

test-unit:
	go test --coverprofile=cover.out --count=1 ./pkg/...

ci-test-unit:
	@go test --json --coverprofile=cover.out --count=1 ./pkg/...
