# Various Scripts in the Key of Golang
Various scripts written in Golang.

## Scripts

### Slack Message Parser
This tool is used to parse the `json`-formatted data that comes from [exporting Slack workspace data](https://slack.com/help/articles/201658943-Export-your-workspace-data).

#### Instructions
1. Move your files to the `SlackMessages` directory.
    ```console
    mkdir -p ./SlackMessages
    cp [SLACK_WORKSPACE_DATA_PATH_HERE]/*.json ./SlackMessages/
    ```
2. Run the Slack message parser using the command: `make parse-slack-data`.
