#!/bin/sh

aws --profile folio-terraform s3 cp s3://production-golio-media/production/folio-db-export/ ./exportdata/ --recursive

pyenv install 3.12.4
pyenv local 3.12.4
python3 -m venv .venv
source .venv/bin/activate

pip3 install -r requirements.txt

python3 ./main.py