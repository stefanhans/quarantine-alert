# quarantine-alert

For the time being using project _gke-serverless-211907_ instead of _quarantine-alert-22365_

```
gcloud iam service-accounts create play-223 \
 --display-name="Play with Firestore" \
 --description="For playing with firestore and eventually more"
```

```
gcloud projects add-iam-policy-binding gke-serverless-211907 \
 --member "serviceAccount:play-223@gke-serverless-211907.iam.gserviceaccount.com" \
 --role "roles/owner"
```

```
gcloud iam service-accounts keys create ../.secrets/play-223-gke-serverless-211907.json \
 --iam-account play-223@gke-serverless-211907.iam.gserviceaccount.com
```

```
export GOOGLE_APPLICATION_CREDENTIALS="/Users/stefan/go/src/github.com/stefanhans/.secrets/play-223-gke-serverless-211907.json"
```

```bash
curl localhost:8080/register -d '{}'
{"reporter":"krMTErfR8qiPCI4FLpnV","contagious":false,"since":"2020-04-01T12:47:01.867622+02:00"}

curl localhost:8080/register -d '{}' 
{"reporter":"NmL2n4u06a69w7CASuND","contagious":false,"since":"2020-04-01T13:04:46.346131+02:00"}
```

```bash

```


gcloud iam service-accounts create server-22365 \
 --display-name="Develop with Firestore" \
 --description="For developing with firestore and eventually more"
 
 gcloud projects add-iam-policy-binding quarantine-alert-22365 \
  --member "serviceAccount:server-22365@quarantine-alert-22365.iam.gserviceaccount.com" \
  --role "roles/owner"
  
  gcloud iam service-accounts keys create ../.secrets/server-22365-quarantine-alert-22365.json \
   --iam-account server-22365@quarantine-alert-22365.iam.gserviceaccount.com
   
export GOOGLE_APPLICATION_CREDENTIALS="/Users/stefan/go/src/github.com/stefanhans/.secrets/server-22365-quarantine-alert-22365.json"