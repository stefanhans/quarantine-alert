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
