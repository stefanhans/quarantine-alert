#### Prepare Deployment of Functions

Initialize Go modules:

```bash
GO111MODULE=on
go mod init && go mod vendor
```

#### Deploy Function ```register```
```bash
gcloud functions deploy register --region europe-west3 \
    --entry-point Register --runtime go113 --trigger-http \
    --allow-unauthenticated
```

#### Test Function ```register```
```bash
gcloud functions describe register --region europe-west3 --format='value(httpsTrigger.url)'

curl https://europe-west3-quarantine-alert-22365.cloudfunctions.net/register -d '{}'
curl https://europe-west3-quarantine-alert-22365.cloudfunctions.net/register -d '{}'
```

#### Deploy Function ```update```
```bash
gcloud functions deploy update --region europe-west3 \
    --entry-point Update --runtime go113 --trigger-http \
    --allow-unauthenticated
```

#### Test Function ```update```
```bash
gcloud functions describe update --region europe-west3 --format='value(httpsTrigger.url)'

curl https://europe-west3-quarantine-alert-22365.cloudfunctions.net/update -d '{"reporter":"<existing ID>","contagious":true}'
```

#### Deploy Function ```contacted```
```bash
gcloud functions deploy contacted --region europe-west3 \
    --entry-point Contacted --runtime go113 --trigger-http \
    --allow-unauthenticated
```

#### Test Function ```contacted```
```bash
gcloud functions describe contacted --region europe-west3 --format='value(httpsTrigger.url)'

curl https://europe-west3-quarantine-alert-22365.cloudfunctions.net/contacted -d '{"reporter":"<1st existing ID>","contact":"<2nd existing ID>"}'
```

#### Deploy Function ```query```
```bash
gcloud functions deploy query --region europe-west3 \
    --entry-point Query --runtime go113 --trigger-http \
    --allow-unauthenticated
```

#### Test Function ```query```
```bash
gcloud functions describe query --region europe-west3 --format='value(httpsTrigger.url)'

curl https://europe-west3-quarantine-alert-22365.cloudfunctions.net/query -d '{"requester":"<1st existing ID>"}'
```

#### Deploy Function ```dump```
```bash
gcloud functions deploy dump --region europe-west3 \
    --entry-point Dump --runtime go113 --trigger-http \
    --allow-unauthenticated
```

#### Test Function ```query```
```bash
gcloud functions describe query --region europe-west3 --format='value(httpsTrigger.url)'

curl https://europe-west3-quarantine-alert-22365.cloudfunctions.net/query -d '{"requester":"<1st existing ID>"}'
```

#### Cleansing

For restoring the functionality of the local development environment, 
and to avoid having two competing Go module structures, 
we have to remove the files.

```bash
rm -r vendor go.mod go.sum
```