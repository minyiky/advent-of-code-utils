test-unit:
	go test --coverage=cover.out --count=1 ./package/...

ci-test-unit:
	go test --coverage=cover.out --count=1 ./package/...