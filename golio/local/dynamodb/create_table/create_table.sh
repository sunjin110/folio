#!/bin/sh

TABLE_NAME="user_sessions"

# すでにテーブルが定義されている場合は作成しない
if aws dynamodb describe-table --table-name $TABLE_NAME --endpoint-url $ENDPOINT_URL > /dev/null 2>&1
then
    echo "テーブル ${TABLE_NAME} はすでに存在します"
    exit 0
fi

aws dynamodb create-table \
    --endpoint-url ${ENDPOINT_URL} \
    --table-name ${TABLE_NAME} \
    --attribute-definitions \
        AttributeName=email,AttributeType=S \
        AttributeName=access_token,AttributeType=S \
    --key-schema \
        AttributeName=email,KeyType=HASH \
    --provisioned-throughput \
        ReadCapacityUnits=3,WriteCapacityUnits=3 \
    --global-secondary-indexes '[
        {
            "IndexName": "access_token_index",
            "KeySchema": [
                {"AttributeName": "access_token", "KeyType": "HASH"}
            ],
            "Projection": {
                "ProjectionType": "ALL"
            },
            "ProvisionedThroughput": {
                "ReadCapacityUnits": 3,
                "WriteCapacityUnits": 3
            }
        }
    ]'
