# CI Watcher Scheduler for OCP QE team.

## How to start using CI-watcher-scheduler:

- [Create a Google Cloud project](https://developers.google.com/workspace/guides/create-project)
- [Enable Google Workspace APIs](https://developers.google.com/workspace/guides/enable-apis)
- [Create access credentials. Service account](https://developers.google.com/workspace/guides/create-credentials#service-account)

Export variables:

```bash
export SPREADSHEET_ID=your_google_docs_spreadsheet_id
export CREDENTIALS=/path/to/your/credentials.json
export DEBUG=true/false # true will add additional debug informations in logs
export E2E_WATCHER_Y0_WEIGHT=17
export E2E_WATCHER_Y1_WEIGHT=16
export E2E_WATCHER_Y2_WEIGHT=10
export E2E_WATCHER_Y3_WEIGHT=8
export E2E_WATCHER_Y4_WEIGHT=3
export E2E_WATCHER_Y5_WEIGHT=2
export UPGR_WATCHER_Y0_WEIGHT=17
export UPGR_WATCHER_Y1_WEIGHT=16
export UPGR_WATCHER_Y2_WEIGHT=10
export UPGR_WATCHER_Y3_WEIGHT=8
export UPGR_WATCHER_Y4_WEIGHT=3
export UPGR_WATCHER_Y5_WEIGHT=2

```

Run scheduler to create a new schedule for next available day.

```bash
go run cmd/scheduler/scheduler.go
```

To run unit test with coverage:

```bash
go test ./... -coverprofile=/tmp/cov.out && go tool cover -html=/tmp/cov.out
```

## TODO (random order)
- Clean up all TODOs in code
- Sending calendar event for Watchers
- getting PTO from calendar
- adding buddy for New To CI Watcher
- getting parameters (like date) from command line
- Adding comments about holidays and PTO in spreadsheet
