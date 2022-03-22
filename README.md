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
export DATE='3/9/2022' # Temportary way of getting day to schedule (format: M/D/YYYY)
```

## TODO (random order)
- Clean up all TODOs in code
- Sending calendar event for Watchers
- getting PTO from calendar
- adding buddy for New To CI Watcher
- getting parameters (like date) from command line
- Adding comments about holidays and PTO in spreadsheet

## Need to decide
- Default trigger? - Jenkins or by user?
- Range of date? - one date at once and if needed loop in jenkins, or everything managed by scheduler?
- Holidays - Take from spreadsheet or from calendar?