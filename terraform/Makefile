
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
	@cd "$(PWD)/env/production" && terraform apply -lock=false --var-file=./secret.tfvars

deploy-init:
	@cd "$(PWD)/env/deployment" && terraform init

deploy-plan:
	@cd "$(PWD)/env/deployment" && terraform plan --var-file=./secret.tfvars

deploy-apply:
	@cd "$(PWD)/env/deployment" && terraform apply --var-file=./secret.tfvars

lint-init:
	@cd "$(PWD)/env/deployment" && tflint --init
	@cd "$(PWD)/env/production" && tflint --init
	@cd "$(PWD)/env/local" && tflint --init

lint:
	cd "$(PWD)/env/deployment" && tflint
	cd "$(PWD)/env/production" && tflint
	cd "$(PWD)/env/local" && tflint

install/pluralith/mac:
	wget https://github.com/Pluralith/pluralith-cli/releases/download/v0.2.2/pluralith_cli_darwin_amd64_v0.2.2
	mv pluralith_cli_darwin_amd64_v0.2.2 pluralith
	chmod +x pluralith
	sudo mv pluralith /usr/local/bin/pluralith
	wget https://github.com/Pluralith/pluralith-cli-graphing-release/releases/download/v0.2.1/pluralith_cli_graphing_darwin_amd64_0.2.1
	mv pluralith_cli_graphing_darwin_amd64_0.2.1 pluralith-cli-graphing
	mkdir -p ~/Pluralith/bin/
	mv pluralith-cli-graphing ~/Pluralith/bin/pluralith-cli-graphing
	chmod +x ~/Pluralith/bin/pluralith-cli-graphing

	pluralith version

generate/terraform/diagram:
	@cd "$(PWD)/env/production" && pluralith graph --local-only --var-file=./secret.tfvars