# quarantine-alert

### Backend

#### Developer

[Developer README](./DEVELOPER.md)

### Cloud Functions


---
Register

```bash
URL: 
https://europe-west3-quarantine-alert-22365.cloudfunctions.net/register

JSON:
reporter                string    : unique ID will be created by Firestore [ignored]
contagious              bool      : set to false as default [optional]
time-contagion-updated  timestamp : set to current timestamp as default [optional]

Examples:
curl https://europe-west3-quarantine-alert-22365.cloudfunctions.net/register -d '{}'

curl https://europe-west3-quarantine-alert-22365.cloudfunctions.net/register -d '{"contagious":false}'

curl https://europe-west3-quarantine-alert-22365.cloudfunctions.net/register -d '{"contagious":true}'

curl https://europe-west3-quarantine-alert-22365.cloudfunctions.net/register \
    -d '{"contagious":true,"time-contagion-updated":"2020-04-01T08:00:00+02:00"}'
    
curl https://europe-west3-quarantine-alert-22365.cloudfunctions.net/register \
    -d '{"contagious":true,"time-contagion-updated":"2020-04-01T00:00:00Z"}'
    
curl https://europe-west3-quarantine-alert-22365.cloudfunctions.net/register \
    -d '{"reporter":"to be ignored","contagious":false,"time-contagion-updated":"2020-04-01T08:46:36.649207+02:00"}'
```

---
Update

```bash
URL: 
https://europe-west3-quarantine-alert-22365.cloudfunctions.net/update

JSON:
reporter    string: 	         : unique ID has to match Firestore's Document ID [mandatory]
contagious  bool                 : to be updated 
time-contagion-updated timestamp : set to current timestamp as default [optional]

Examples:
curl https://europe-west3-quarantine-alert-22365.cloudfunctions.net/update \
    -d '{"reporter":"B3AQeB8acvRzK9FXtPxP","contagious":true}'

```

---

Contacted

```bash
URL: 
https://europe-west3-quarantine-alert-22365.cloudfunctions.net/contacted

JSON:
reporter     string     : unique ID of reporting app has to match Firestore's Document ID
contact      string     : unique ID of contacted app
contact-time timestamp  : set to current timestamp as default [optional]

Examples:
curl https://europe-west3-quarantine-alert-22365.cloudfunctions.net/contacted \
    -d '{"reporter":"B3AQeB8acvRzK9FXtPxP","contact":"HlUv5MfjBhbfvEtNfATR"}'
```

---
Query

```bash
URL: 
https://europe-west3-quarantine-alert-22365.cloudfunctions.net/query

JSON:
requester               string    : unique ID of requester
contagious              bool      : set to false as default [ignored]
time-contagion-updated  timestamp : set to current timestamp as default [ignored]

Examples:
curl https://europe-west3-quarantine-alert-22365.cloudfunctions.net/query -d '{"requester":"uBAJbYDJTBHOVqrceZur"}'
```

---
Dump

```bash
URL: 
https://europe-west3-quarantine-alert-22365.cloudfunctions.net/dump

JSON:
indent string : JSON data formatted with indent [default: false]

Examples:
curl https://europe-west3-quarantine-alert-22365.cloudfunctions.net/dump -d '{}'
curl https://europe-west3-quarantine-alert-22365.cloudfunctions.net/dump -d '{"indent":true}'
curl https://europe-west3-quarantine-alert-22365.cloudfunctions.net/dump -d '{"indent":false}'
```



