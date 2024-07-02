#!/bin/sh

# すでにテーブルが定義されている場合は作成しない
if aws dynamodb describe-table --table-name user_sessions --endpoint-url $ENDPOINT_URL > /dev/null 2>&1
then
    echo "テーブル user_sessions はすでに存在します"
else
aws dynamodb create-table \
    --endpoint-url ${ENDPOINT_URL} \
    --table-name user_sessions \
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
fi

if aws dynamodb describe-table --table-name user_sessions_v2 --endpoint-url $ENDPOINT_URL > /dev/null 2>&1
then
    echo "テーブル user_sessions_v2はすでに存在します"
else
aws dynamodb create-table \
    --endpoint-url ${ENDPOINT_URL} \
    --table-name user_sessions_v2 \
    --attribute-definitions \
        AttributeName=access_token,AttributeType=S \
        AttributeName=email,AttributeType=S \
        AttributeName=expire_time,AttributeType=N \
    --key-schema \
        AttributeName=access_token,KeyType=HASH \
    --billing-mode PAY_PER_REQUEST \
    --global-secondary-indexes \
        "[
            {
                \"IndexName\": \"email_access_token_index\",
                \"KeySchema\": [
                    {\"AttributeName\":\"email\", \"KeyType\":\"HASH\"},
                    {\"AttributeName\":\"access_token\", \"KeyType\":\"RANGE\"}
                ],
                \"Projection\": {
                    \"ProjectionType\":\"ALL\"
                }
            }
        ]" \
    --tags \
        Key=Name,Value=user_sessions_v2

aws dynamodb update-time-to-live \
    --endpoint-url ${ENDPOINT_URL} \
    --table-name user_sessions_v2 \
    --time-to-live-specification "Enabled=true, AttributeName=expire_time"

fi

aws dynamodb describe-time-to-live --endpoint-url $ENDPOINT_URL --table-name user_sessions_v2

if aws dynamodb describe-table --table-name users --endpoint-url $ENDPOINT_URL > /dev/null 2>&1
then
    echo "テーブル usersはすでに存在します"
else
aws dynamodb create-table \
    --endpoint-url ${ENDPOINT_URL} \
    --table-name users \
    --attribute-definitions \
        AttributeName=email,AttributeType=S \
    --key-schema \
        AttributeName=email,KeyType=HASH \
    --billing-mode PAY_PER_REQUEST \
    --tags \
        Key=Name,Value=users
fi
