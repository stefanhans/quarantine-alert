# quarantine-alert

### Backend

#### Developer

[Developer README](./DEVELOPER.md)
  
### API 


### Cloud Functions

Register

```bash
URL: 
https://europe-west3-quarantine-alert-22365.cloudfunctions.net/register

JSON:
reporter    string: 	: unique ID will be created by Firestore [ignored]
contagious  bool        : set to false as default [optional]
since       timestamp   : set to current timestamp as default [optional]

Examples:
curl https://europe-west3-quarantine-alert-22365.cloudfunctions.net/register -d '{}'

curl https://europe-west3-quarantine-alert-22365.cloudfunctions.net/register -d '{"contagious":false}'

curl https://europe-west3-quarantine-alert-22365.cloudfunctions.net/register -d '{"contagious":true}'

curl https://europe-west3-quarantine-alert-22365.cloudfunctions.net/register -d '{"contagious":true,"since":"2020-04-01T08:00:00+02:00"}'
    
curl https://europe-west3-quarantine-alert-22365.cloudfunctions.net/register -d '{"contagious":true,"since":"2020-04-01T00:00:00Z"}'
    
curl https://europe-west3-quarantine-alert-22365.cloudfunctions.net/register -d '{"reporter":"to be ignored","contagious":false,"since":"2020-04-01T08:46:36.649207+02:00"}'
```

Register

```bash
URL: 
https://europe-west3-quarantine-alert-22365.cloudfunctions.net/register

JSON:
reporter    string: 	: unique ID will be created by Firestore [ignored]
contagious  bool        : set to false as default [optional]
since       timestamp   : set to current timestamp as default [optional]

Examples:
curl https://europe-west3-quarantine-alert-22365.cloudfunctions.net/register -d '{}'

curl https://europe-west3-quarantine-alert-22365.cloudfunctions.net/register -d '{"contagious":false}'

curl https://europe-west3-quarantine-alert-22365.cloudfunctions.net/register -d '{"contagious":true}'

curl https://europe-west3-quarantine-alert-22365.cloudfunctions.net/register -d '{"contagious":true,"since":"2020-04-01T08:00:00+02:00"}'
    
curl https://europe-west3-quarantine-alert-22365.cloudfunctions.net/register -d '{"contagious":true,"since":"2020-04-01T00:00:00Z"}'
    
curl https://europe-west3-quarantine-alert-22365.cloudfunctions.net/register -d '{"reporter":"to be ignored","contagious":false,"since":"2020-04-01T08:46:36.649207+02:00"}'
```



