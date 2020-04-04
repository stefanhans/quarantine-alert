The go client is mainly for testing, for the time being.

#### Testing against localhost
Start the local server in one terminal
```bash
cd ../server
go build main.go && ./main
```

Do the testing in another terminal
```bash
GCP_SERVICE_URL="http://localhost:8080"

go test
```

#### Testing against Cloud Functions
```bash
GCP_SERVICE_URL="https://europe-west3-quarantine-alert-22365.cloudfunctions.net"

go test
```