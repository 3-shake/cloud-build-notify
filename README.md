# Cloud Build Notify

## Getting Started
1. GCP / Bucket / Slack それぞれ必須項目を terraform/terraform.tfvars に設定
2. cloud function にdeploy
```
make terraform.init
make terraform.apply
```
