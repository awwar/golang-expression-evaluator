lint_setup:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.2.0


fmt:
	golangci-lint fmt

fix:
	golangci-lint run --fix