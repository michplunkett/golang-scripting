# Various Scripts in the Key of Golang
Various scripts written in Golang.

## Scripts

### Slack Message Parser
This tool is used to parse the `json`-formatted data that comes from [exporting Slack workspace data](https://slack.com/help/articles/201658943-Export-your-workspace-data).

#### Instructions
```console
mkdir -p ./SlackMessages
cp [YOUR_ARCHIVE_PATH_HERE]/*.json ./SlackMessages/
go run ./cmd/slackMessageParser/main.go
```
