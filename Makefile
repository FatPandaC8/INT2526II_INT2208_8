.PHONY: test test-bank test-triangle coverage clean

# Run all tests
test:
	go test ./...

# Run only bank tests
test-bank:
	go test ./internal/bank -v

# Run only triangle tests
test-triangle:
	go test ./internal/triangle -v

# Run tests with coverage
cover:
	go test ./... -cover

# Generate detailed coverage report
coverage-html:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

# Clean generated files
clean:
	rm -f coverage.out