default: lint

.PHONY: init
init:
	go mod download

.PHONY: lint
lint:
	go mod tidy
	go fmt ./...
	go vet ./...
	staticcheck ./...

.PHONY: parse-slack-data
parse-slack-data:
	go run ./cmd/slackMessageParser/main.go
