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

.PHONY: delete-reddit-history
delete-reddit-history:
	go run ./cmd/deleteRedditHistory/main.go

.PHONY: parse-slack-data
parse-slack-data:
	go run ./cmd/slackMessageParser/main.go
