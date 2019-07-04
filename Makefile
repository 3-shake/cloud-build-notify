terraform.init:
	-@gsutil mb gs://cloud-build-notify-terraform
	@cd ./terraform && terraform init -backend-config ./backend.tfvars  .

terraform.plan:
	@cd ./terraform && terraform plan .

terraform.apply:
	@cd ./terraform && terraform apply .

terraform.destroy:
	@cd ./terraform && terraform destroy .
