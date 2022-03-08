# CI Watcher Scheduler for OCP QE team.

## How to start using CI-watcher-scheduler:

- [Create a Google Cloud project](https://developers.google.com/workspace/guides/create-project)
- [Enable Google Workspace APIs](https://developers.google.com/workspace/guides/enable-apis)
- [Create access credentials. Service account](https://developers.google.com/workspace/guides/create-credentials#service-account)

Export variables

```bash
export SPREADSHEET_ID=your_google_docs_spreadsheet_id
export CREDENTIALS=/path/to/your/credentials.json
export DEBUG=true/false # true will add additional debug informations in logs
export CI=true/false # true will turn off colors from logs
```