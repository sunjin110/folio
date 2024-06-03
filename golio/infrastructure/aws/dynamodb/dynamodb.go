package dynamodb

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// 参考: https://github.com/awsdocs/aws-doc-sdk-examples/blob/main/gov2/dynamodb/actions/table_basics.go

// ErrNotFound 対象のデータが見つかりませんでした
var ErrNotFound = errors.New("not found")

const (
	// BatchGetItemで一度に取得できるデータの最大数
	// https://docs.aws.amazon.com/ja_jp/amazondynamodb/latest/developerguide/ServiceQuotas.html#limits-api
	BatchGetItemMaxItemCount = 100

	// BatchWriteItemで一度に作成または削除できるデータの最大数
	// https://docs.aws.amazon.com/pdfs/amazondynamodb/latest/APIReference/dynamodb-api.pdf#API_BatchWriteItem
	BatchWriteItemMaxItemCount = 25
)

// Client dynamodb内のデータを操作するclient。利用方法はclient_test.goを参考にしてください
// pagingのtokenを*stringで返しているわけ
// - AWS SDKのmap[string]types.AttributeValueをdomainロジックに持たせたくない
// - 上記のkeyをinterfaceとして抽象化して扱うことも考えたが、handlerから渡すstructとdynamoから返すstructが変わるため具象化が複雑になる
// - dynamodb/client.goでpagingのtoken生成を完結させることができるため`*string`を選択しました
//
//go:generate mockgen -source ./client.go -destination ./mock/client.go -package mock_dynamodb
type Client[T DynamoDTO] interface {
	// Add 対象のデータを追加します
	Add(ctx context.Context, table string, dto T) error

	// Update 対象のデータをUpdateします
	Update(ctx context.Context, table string, key map[string]types.AttributeValue, update expression.UpdateBuilder) error

	// Get 対象のデータを取得する
	// Error: ErrNotFound
	Get(ctx context.Context, table string, key map[string]types.AttributeValue) (*T, error)

	// Delete 対象のデータを削除します、対象データが存在しない場合はエラーは返しません
	Delete(ctx context.Context, table string, key map[string]types.AttributeValue) error

	// Query Queryで検索
	Query(ctx context.Context, table string, expr expression.Expression, limit *int32, indexName *string, startEvaluatedKey *string, scanIndexForward *bool) (list []*T, lastEvaluatedKey *string, err error)

	// Scan フルスキャンで検索
	Scan(ctx context.Context, table string, expr expression.Expression, limit *int32, startEvaluatedKey *string) (list []*T, lastEvaluatedKey *string, err error)

	// BatchGetItem 複数のGetItemを一度に投げることができる、tableをまたぐ処理なので軽くwrapだけしている
	BatchGetItem(ctx context.Context, input *BatchGetItemInput) (*BatchGetItemOutput, error)

	// BatchWriteItem 複数のWriteItemを一度に投げることができる、tableを跨ぐ処理なので軽くwrapだけしている
	BatchWriteItem(ctx context.Context, input *BatchWriteItemInput) (*BatchWriteItemOutput, error)
}

type client[T DynamoDTO] struct {
	dynamoDBClient *dynamodb.Client
}

func NewClient[T DynamoDTO](dynamoDBClient *dynamodb.Client) Client[T] {
	return &client[T]{
		dynamoDBClient: dynamoDBClient,
	}
}

func NewInnerClient(cfg aws.Config) *dynamodb.Client {
	return dynamodb.NewFromConfig(cfg)
}

type DynamoDTO interface {
	// IsDynamoDTO domain modelを追加しないようにするためのもの
	IsDynamoDTO()

	// GetKey GetItemで利用できるkeyを作成
	GetKey() (map[string]types.AttributeValue, error)
}

func (client *client[T]) Add(ctx context.Context, table string, dto T) error {
	item, err := attributevalue.MarshalMap(dto)
	if err != nil {
		return fmt.Errorf("failed Marshal to dynamo attribute value. %w", err)
	}
	if _, err := client.dynamoDBClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(table),
		Item:      item,
	}); err != nil {
		return fmt.Errorf("failed PutItem to dynamodb. table: %s, err: %w", table, err)
	}
	return nil
}

func (client *client[T]) Update(ctx context.Context, table string, key map[string]types.AttributeValue, update expression.UpdateBuilder) error {
	expr, err := expression.NewBuilder().WithUpdate(update).Build()
	if err != nil {
		return fmt.Errorf("failed build expression for update. %w", err)
	}

	if _, err := client.dynamoDBClient.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName:                 aws.String(table),
		Key:                       key,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
		ReturnValues:              types.ReturnValueNone, // 必要な場合は返すようにする
	}); err != nil {
		return fmt.Errorf("failed UpdateItem. table: %s, err: %w", table, err)
	}
	return nil
}

// Get outはpointerである必要があります
func (client *client[T]) Get(ctx context.Context, table string, key map[string]types.AttributeValue) (*T, error) {
	response, err := client.dynamoDBClient.GetItem(ctx, &dynamodb.GetItemInput{
		Key:       key,
		TableName: aws.String(table),
	})

	if err != nil {
		return nil, fmt.Errorf("failed GetItem. key: %+v, err: %w", key, err)
	}

	if len(response.Item) == 0 {
		return nil, ErrNotFound
	}

	out := new(T)
	if err := attributevalue.UnmarshalMap(response.Item, out); err != nil {
		return nil, fmt.Errorf("failed GetItem Unmarshal response. key: %+v, res: %+v, err: %w", key, response, err)
	}
	return out, nil
}

type BatchGetItemInput = dynamodb.BatchGetItemInput
type BatchGetItemOutput = dynamodb.BatchGetItemOutput
type BatchWriteItemInput = dynamodb.BatchWriteItemInput
type BatchWriteItemOutput = dynamodb.BatchWriteItemOutput

func (client *client[T]) BatchGetItem(ctx context.Context, input *BatchGetItemInput) (*BatchGetItemOutput, error) {
	output, err := client.dynamoDBClient.BatchGetItem(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed GetBatchItem. %w", err)
	}
	return output, nil
}

func (client *client[T]) BatchWriteItem(ctx context.Context, input *BatchWriteItemInput) (*BatchWriteItemOutput, error) {
	requestCount := 0
	for _, writeRequests := range input.RequestItems {
		requestCount += len(writeRequests)
	}

	if requestCount > BatchWriteItemMaxItemCount {
		return nil, fmt.Errorf("requestItems are less than %d items", BatchWriteItemMaxItemCount+1)
	}
	output, err := client.dynamoDBClient.BatchWriteItem(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed BatchWriteItem: %w", err)
	}
	return output, nil
}

func (client *client[T]) Delete(ctx context.Context, table string, key map[string]types.AttributeValue) error {
	if _, err := client.dynamoDBClient.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(table),
		Key:       key,
	}); err != nil {
		return fmt.Errorf("failed delete item from dynamodb. %w", err)
	}
	return nil
}

func (client *client[T]) Query(ctx context.Context, table string, expr expression.Expression, limit *int32, indexName *string, startEvaluatedKey *string, scanIndexForword *bool) ([]*T, *string, error) {
	exclusiveStartKey, err := client.decodeStartEvaluatedKey(startEvaluatedKey)
	if err != nil {
		return nil, nil, fmt.Errorf("failed decode startEvaluatedKey of query. %w", err)
	}

	response, err := client.dynamoDBClient.Query(ctx, &dynamodb.QueryInput{
		TableName:                 aws.String(table),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		ProjectionExpression:      expr.Projection(),
		ExclusiveStartKey:         exclusiveStartKey,
		Limit:                     limit,
		ScanIndexForward:          scanIndexForword,
		IndexName:                 indexName,
	})

	if err != nil {
		return nil, nil, fmt.Errorf("failed dynamodb query. %w", err)
	}

	lastEvaluatedKey, err := client.encodeLastEvaluatedKey(response.LastEvaluatedKey)
	if err != nil {
		return nil, nil, fmt.Errorf("failed encode last evaluated key of query. %w", err)
	}

	dtos := make([]*T, 0)
	if err := attributevalue.UnmarshalListOfMaps(response.Items, &dtos); err != nil {
		return nil, nil, fmt.Errorf("failed unmarshal query response. %w", err)
	}
	return dtos, lastEvaluatedKey, nil
}

func (client *client[T]) Scan(ctx context.Context, table string, expr expression.Expression, limit *int32, startEvaluatedKey *string) ([]*T, *string, error) {
	exclusiveStartKey, err := client.decodeStartEvaluatedKey(startEvaluatedKey)
	if err != nil {
		return nil, nil, fmt.Errorf("failed decode startEvaluatedKey of scan. %w", err)
	}

	response, err := client.dynamoDBClient.Scan(ctx, &dynamodb.ScanInput{
		TableName:                 aws.String(table),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		// filterはフルスキャンデータをfilterするだけなので、条件によって0件で帰ってくるがページングできるパターンもあります
		FilterExpression:  expr.Filter(),
		ExclusiveStartKey: exclusiveStartKey,
		Limit:             limit,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed dynamodb scan. %w", err)
	}

	lastEvaluatedKey, err := client.encodeLastEvaluatedKey(response.LastEvaluatedKey)
	if err != nil {
		return nil, nil, fmt.Errorf("failed encode last evaluated key of scan. %w", err)
	}

	dtos := make([]*T, 0)
	if err := attributevalue.UnmarshalListOfMaps(response.Items, &dtos); err != nil {
		return nil, nil, fmt.Errorf("failed unmarshal scan response. %w", err)
	}
	return dtos, lastEvaluatedKey, nil
}

func (*client[T]) decodeStartEvaluatedKey(hashKey *string) (map[string]types.AttributeValue, error) {
	if hashKey == nil {
		return nil, nil
	}

	b, err := base64.StdEncoding.DecodeString(*hashKey)
	if err != nil {
		return nil, fmt.Errorf("failed decode startEvaluatedKey. %w", err)
	}

	t := &map[string]interface{}{}
	if err := json.Unmarshal(b, t); err != nil {
		return nil, fmt.Errorf("failed unmarshal startEvaluatedKey. key: %s, err: %w", string(b), err)
	}

	key, err := attributevalue.MarshalMap(t)
	if err != nil {
		return nil, fmt.Errorf("failed marshal startEvaluatedKey. key: %s, err: %w", string(b), err)
	}
	return key, nil
}

func (*client[T]) encodeLastEvaluatedKey(key map[string]types.AttributeValue) (*string, error) {
	if len(key) == 0 {
		return nil, nil
	}

	t := &map[string]interface{}{}
	if err := attributevalue.UnmarshalMap(key, t); err != nil {
		return nil, fmt.Errorf("failed unmarshal lastEvaluatedKey. err: %w", err)
	}
	b, err := json.Marshal(t)
	if err != nil {
		return nil, fmt.Errorf("failed json marshal lastEvaluatedKey. key: %+v, err: %w", t, err)
	}
	hash := base64.StdEncoding.EncodeToString(b)
	return &hash, nil
}
