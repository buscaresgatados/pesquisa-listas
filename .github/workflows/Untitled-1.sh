REVISIONS=$(gcloud run revisions list --service=refugio-rs-stg --region=southamerica-east1 --format=json)
PREVIOUS_REVISION=$(jq -r '.[1].metadata.name' <<< ${REVISIONS}) 
NEW_REVISION=$(jq -r '.[0].metadata.name' <<< ${REVISIONS}) 
gcloud run services update-traffic refugio-rs-stg --region=southamerica-east1 --to-revisions=$PREVIOUS_REVISION=100,$NEW_REVISION=0