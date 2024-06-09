
# google credentailsのログ

```sh
gcloud iam service-accounts create folio-terraform --display-name "folio terraform"

gcloud projects add-iam-policy-binding folio-sunjin\
  --member "serviceAccount:folio-terraform@folio-sunjin.iam.gserviceaccount.com"\
  --role "roles/owner"

gcloud iam service-accounts keys create credentials-terraform.json --iam-account folio-terraform@folio-sunjin.iam.gserviceaccount.com


```