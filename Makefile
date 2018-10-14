linux: ## Creates linux binary to bin/seat-distributor-linux-amd64
	GOOS=linux GOARCH=amd64 go build -o="bin/seat-distributor-linux-amd64"

osx: ## Creates osx binary to bin/seat-distributor-darwin-amd64
	GOOS=darwin GOARCH=amd64 go build -o="bin/seat-distributor-darwin-amd64"

windows: ## Creates windows binary to bin/seat-distributor-windows-amd64.exe
	GOOS=windows GOARCH=amd64 go build -o="bin/seat-distributor-windows-amd64.exe"

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help