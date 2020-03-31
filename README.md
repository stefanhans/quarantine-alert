# quarantine-alert

```
gcloud iam service-accounts create play-223 \
 --display-name="Play with Firestore" \
 --description="For playing with firestore and eventually more"
```

```
gcloud projects add-iam-policy-binding quarantine-alert-22365 \
 --member "serviceAccount:play-223@quarantine-alert-22365.iam.gserviceaccount.com" \
 --role "roles/owner"
```

```
gcloud iam service-accounts keys create ../.secrets/play-223-quarantine-alert-22365.json \
 --iam-account play-223@quarantine-alert-22365.iam.gserviceaccount.com
```
