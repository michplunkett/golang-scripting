# Various Scripts in the Key of Golang
Various scripts written in Golang.

## Scripts

### Reddit Comment and Post Deleter
This is used to delete all comments for a given user that are older than 2-years-old. To properly configure, please 
create the following environment variables: `REDDIT_USER_ID`, `REDDIT_USER_PASSWORD`, `REDDIT_APP_ID`, `REDDIT_SECRET`.

For information regarding the acquisition of values for the above environment variables, please 
check here: [Link](https://github.com/reddit-archive/reddit/wiki/OAuth2-Quick-Start-Example#first-steps).

### Slack Message Parser
This script is used to parse the `json`-formatted data that comes from 
[exporting Slack workspace data](https://slack.com/help/articles/201658943-Export-your-workspace-data) into a `csv` 
file named: `slack_records.csv`.

| TimeStamp | UserID | UserName | RealName | MessageType |  Text  | Attachments |  Files   |
|:---------:|:------:|:--------:|:--------:|:-----------:|:------:|:-----------:|:--------:|
|  string   | string |  string  |  string  |   string    | string |  []string   | []string |

#### Instructions
1. Move your files to the `SlackMessages` directory.
    ```console
    mkdir -p ./SlackMessages
    cp [SLACK_WORKSPACE_DATA_PATH_HERE]/*.json ./SlackMessages/
    ```
2. Run the Slack message parser using the command: `make parse-slack-data`.
