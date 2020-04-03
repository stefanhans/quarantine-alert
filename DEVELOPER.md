# quarantine-alert for developer

### Environment for local development

These commandline snippets show how to initialize your environment, if you have already a GCP project created 
with enabled Firestore component.

```bash
export GCP_PROJECT="<google cloud project-id>"
export CREDENTIALS_DIR="<local directory to store credentials>" 
export GCP_SERVICE_ACCOUNT="<name of service account>"
```

e.g.
```bash
export GCP_PROJECT="quarantine-alert-22365"
export CREDENTIALS_DIR="$HOME/go/src/github.com/stefanhans/.secrets"
export GCP_SERVICE_ACCOUNT="server-22365"
```

```bash
gcloud iam service-accounts create ${GCP_SERVICE_ACCOUNT} \
 --display-name="Develop with Firestore" \
 --description="For developing with firestore and eventually more"
```

```bash
gcloud projects add-iam-policy-binding ${GCP_PROJECT} \
--member "serviceAccount:${GCP_SERVICE_ACCOUNT}@${GCP_PROJECT}.iam.gserviceaccount.com" \
--role "roles/owner"
```

```bash
gcloud iam service-accounts keys create ${CREDENTIALS_DIR}/${GCP_SERVICE_ACCOUNT}-${GCP_PROJECT}.json \
--iam-account ${GCP_SERVICE_ACCOUNT}@${GCP_PROJECT}.iam.gserviceaccount.com
```

```bash
export GOOGLE_APPLICATION_CREDENTIALS="${CREDENTIALS_DIR}/${GCP_SERVICE_ACCOUNT}-quarantine-alert-22365.json"
```


 

  

