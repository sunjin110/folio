
stg-init:
	@cd "$(PWD)/env/staging" && terraform init

local-init:
	@cd "$(PWD)/env/local" && terraform init

local-plan:
	@cd "$(PWD)/env/local" && terraform plan --var-file=./secret.tfvars

local-apply:
	@cd "$(PWD)/env/local" && terraform apply --var-file=./secret.tfvars

fmt:
	terraform fmt -recursive

prd-init:
	@cd "$(PWD)/env/production" && terraform init

prd-plan:
	@cd "$(PWD)/env/production" && terraform plan --var-file=./secret.tfvars

prd-apply:
	@cd "$(PWD)/env/production" && terraform apply --var-file=./secret.tfvars
