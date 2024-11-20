# \GolioAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArticleTagsGet**](GolioAPI.md#ArticleTagsGet) | **Get** /article_tags | 記事タグ一覧取得
[**ArticleTagsPost**](GolioAPI.md#ArticleTagsPost) | **Post** /article_tags | 記事タグの作成
[**ArticleTagsTagIdDelete**](GolioAPI.md#ArticleTagsTagIdDelete) | **Delete** /article_tags/{tag_id} | 記事タグの削除
[**ArticleTagsTagIdPut**](GolioAPI.md#ArticleTagsTagIdPut) | **Put** /article_tags/{tag_id} | 記事タグの更新
[**ArticlesAiPost**](GolioAPI.md#ArticlesAiPost) | **Post** /articles/ai | 記事AI作成
[**ArticlesArticleIdAiPut**](GolioAPI.md#ArticlesArticleIdAiPut) | **Put** /articles/{article_id}/ai | 記事AI更新
[**ArticlesArticleIdGet**](GolioAPI.md#ArticlesArticleIdGet) | **Get** /articles/{article_id} | 記事取得
[**ArticlesArticleIdPut**](GolioAPI.md#ArticlesArticleIdPut) | **Put** /articles/{article_id} | 記事更新
[**ArticlesGet**](GolioAPI.md#ArticlesGet) | **Get** /articles | 記事一覧取得
[**ArticlesPost**](GolioAPI.md#ArticlesPost) | **Post** /articles | 記事投稿
[**EnglishDictionaryWordBookmarkDelete**](GolioAPI.md#EnglishDictionaryWordBookmarkDelete) | **Delete** /english_dictionary/{word}/bookmark | 辞書で引いた単語ブックマークを削除
[**EnglishDictionaryWordBookmarkPut**](GolioAPI.md#EnglishDictionaryWordBookmarkPut) | **Put** /english_dictionary/{word}/bookmark | 辞書で引いた単語ブックマークを作成
[**EnglishDictionaryWordGet**](GolioAPI.md#EnglishDictionaryWordGet) | **Get** /english_dictionary/{word} | 英単語を辞書で引く
[**HelloGet**](GolioAPI.md#HelloGet) | **Get** /hello | hello
[**MediaGet**](GolioAPI.md#MediaGet) | **Get** /media | メディア一覧取得
[**MediaMediumIdDelete**](GolioAPI.md#MediaMediumIdDelete) | **Delete** /media/{medium_id} | メディアの削除
[**MediaMediumIdGet**](GolioAPI.md#MediaMediumIdGet) | **Get** /media/{medium_id} | メディアの取得
[**MediaPost**](GolioAPI.md#MediaPost) | **Post** /media | メディアの登録
[**TasksIdGet**](GolioAPI.md#TasksIdGet) | **Get** /tasks/{id} | タスク詳細取得
[**TasksIdPut**](GolioAPI.md#TasksIdPut) | **Put** /tasks/{id} | タスク編集
[**TasksPost**](GolioAPI.md#TasksPost) | **Post** /tasks | タスク作成
[**TranslationPost**](GolioAPI.md#TranslationPost) | **Post** /translation | 翻訳



## ArticleTagsGet

> ArticleTagsGet200Response ArticleTagsGet(ctx).SearchText(searchText).Offset(offset).Limit(limit).Execute()

記事タグ一覧取得



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sunjin110/github.com/sunjin110/folio/lime"
)

func main() {
	searchText := "searchText_example" // string |  (optional)
	offset := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GolioAPI.ArticleTagsGet(context.Background()).SearchText(searchText).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GolioAPI.ArticleTagsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArticleTagsGet`: ArticleTagsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `GolioAPI.ArticleTagsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArticleTagsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchText** | **string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**ArticleTagsGet200Response**](ArticleTagsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArticleTagsPost

> InsertArticleTagResponse ArticleTagsPost(ctx).ArticleTagsPostRequest(articleTagsPostRequest).Execute()

記事タグの作成



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sunjin110/github.com/sunjin110/folio/lime"
)

func main() {
	articleTagsPostRequest := *openapiclient.NewArticleTagsPostRequest("Name_example") // ArticleTagsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GolioAPI.ArticleTagsPost(context.Background()).ArticleTagsPostRequest(articleTagsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GolioAPI.ArticleTagsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArticleTagsPost`: InsertArticleTagResponse
	fmt.Fprintf(os.Stdout, "Response from `GolioAPI.ArticleTagsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArticleTagsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **articleTagsPostRequest** | [**ArticleTagsPostRequest**](ArticleTagsPostRequest.md) |  | 

### Return type

[**InsertArticleTagResponse**](InsertArticleTagResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArticleTagsTagIdDelete

> DeleteArticleTagResponse ArticleTagsTagIdDelete(ctx, tagId).Execute()

記事タグの削除



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sunjin110/github.com/sunjin110/folio/lime"
)

func main() {
	tagId := "tagId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GolioAPI.ArticleTagsTagIdDelete(context.Background(), tagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GolioAPI.ArticleTagsTagIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArticleTagsTagIdDelete`: DeleteArticleTagResponse
	fmt.Fprintf(os.Stdout, "Response from `GolioAPI.ArticleTagsTagIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiArticleTagsTagIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteArticleTagResponse**](DeleteArticleTagResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArticleTagsTagIdPut

> UpdateArticleTagResponse ArticleTagsTagIdPut(ctx, tagId).ArticleTagsTagIdPutRequest(articleTagsTagIdPutRequest).Execute()

記事タグの更新



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sunjin110/github.com/sunjin110/folio/lime"
)

func main() {
	tagId := "tagId_example" // string | 
	articleTagsTagIdPutRequest := *openapiclient.NewArticleTagsTagIdPutRequest("Name_example") // ArticleTagsTagIdPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GolioAPI.ArticleTagsTagIdPut(context.Background(), tagId).ArticleTagsTagIdPutRequest(articleTagsTagIdPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GolioAPI.ArticleTagsTagIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArticleTagsTagIdPut`: UpdateArticleTagResponse
	fmt.Fprintf(os.Stdout, "Response from `GolioAPI.ArticleTagsTagIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiArticleTagsTagIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **articleTagsTagIdPutRequest** | [**ArticleTagsTagIdPutRequest**](ArticleTagsTagIdPutRequest.md) |  | 

### Return type

[**UpdateArticleTagResponse**](UpdateArticleTagResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArticlesAiPost

> ArticlesAiPost200Response ArticlesAiPost(ctx).ArticlesAiPostRequest(articlesAiPostRequest).Execute()

記事AI作成



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sunjin110/github.com/sunjin110/folio/lime"
)

func main() {
	articlesAiPostRequest := *openapiclient.NewArticlesAiPostRequest("Prompt_example") // ArticlesAiPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GolioAPI.ArticlesAiPost(context.Background()).ArticlesAiPostRequest(articlesAiPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GolioAPI.ArticlesAiPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArticlesAiPost`: ArticlesAiPost200Response
	fmt.Fprintf(os.Stdout, "Response from `GolioAPI.ArticlesAiPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArticlesAiPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **articlesAiPostRequest** | [**ArticlesAiPostRequest**](ArticlesAiPostRequest.md) |  | 

### Return type

[**ArticlesAiPost200Response**](ArticlesAiPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArticlesArticleIdAiPut

> ArticlesArticleIdAiPut200Response ArticlesArticleIdAiPut(ctx, articleId).ArticlesArticleIdAiPutRequest(articlesArticleIdAiPutRequest).Execute()

記事AI更新



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sunjin110/github.com/sunjin110/folio/lime"
)

func main() {
	articleId := "articleId_example" // string | 
	articlesArticleIdAiPutRequest := *openapiclient.NewArticlesArticleIdAiPutRequest("Message_example") // ArticlesArticleIdAiPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GolioAPI.ArticlesArticleIdAiPut(context.Background(), articleId).ArticlesArticleIdAiPutRequest(articlesArticleIdAiPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GolioAPI.ArticlesArticleIdAiPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArticlesArticleIdAiPut`: ArticlesArticleIdAiPut200Response
	fmt.Fprintf(os.Stdout, "Response from `GolioAPI.ArticlesArticleIdAiPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiArticlesArticleIdAiPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **articlesArticleIdAiPutRequest** | [**ArticlesArticleIdAiPutRequest**](ArticlesArticleIdAiPutRequest.md) |  | 

### Return type

[**ArticlesArticleIdAiPut200Response**](ArticlesArticleIdAiPut200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArticlesArticleIdGet

> ArticlesArticleIdGet200Response ArticlesArticleIdGet(ctx, articleId).Execute()

記事取得



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sunjin110/github.com/sunjin110/folio/lime"
)

func main() {
	articleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GolioAPI.ArticlesArticleIdGet(context.Background(), articleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GolioAPI.ArticlesArticleIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArticlesArticleIdGet`: ArticlesArticleIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `GolioAPI.ArticlesArticleIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiArticlesArticleIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ArticlesArticleIdGet200Response**](ArticlesArticleIdGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArticlesArticleIdPut

> map[string]interface{} ArticlesArticleIdPut(ctx, articleId).ArticlesPostRequest(articlesPostRequest).Execute()

記事更新



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sunjin110/github.com/sunjin110/folio/lime"
)

func main() {
	articleId := "articleId_example" // string | 
	articlesPostRequest := *openapiclient.NewArticlesPostRequest("Title_example", "Body_example", []string{"TagIds_example"}) // ArticlesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GolioAPI.ArticlesArticleIdPut(context.Background(), articleId).ArticlesPostRequest(articlesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GolioAPI.ArticlesArticleIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArticlesArticleIdPut`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GolioAPI.ArticlesArticleIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiArticlesArticleIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **articlesPostRequest** | [**ArticlesPostRequest**](ArticlesPostRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArticlesGet

> ArticlesGet200Response ArticlesGet(ctx).Offset(offset).Limit(limit).SearchTitleText(searchTitleText).Tags(tags).Execute()

記事一覧取得



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sunjin110/github.com/sunjin110/folio/lime"
)

func main() {
	offset := int32(0) // int32 |  (optional)
	limit := int32(10) // int32 |  (optional)
	searchTitleText := "searchTitleText_example" // string | タイトルの検索 (optional)
	tags := []string{"Inner_example"} // []string | tagのarray (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GolioAPI.ArticlesGet(context.Background()).Offset(offset).Limit(limit).SearchTitleText(searchTitleText).Tags(tags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GolioAPI.ArticlesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArticlesGet`: ArticlesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `GolioAPI.ArticlesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArticlesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **searchTitleText** | **string** | タイトルの検索 | 
 **tags** | **[]string** | tagのarray | 

### Return type

[**ArticlesGet200Response**](ArticlesGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArticlesPost

> ArticlesPost200Response ArticlesPost(ctx).ArticlesPostRequest(articlesPostRequest).Execute()

記事投稿



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sunjin110/github.com/sunjin110/folio/lime"
)

func main() {
	articlesPostRequest := *openapiclient.NewArticlesPostRequest("Title_example", "Body_example", []string{"TagIds_example"}) // ArticlesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GolioAPI.ArticlesPost(context.Background()).ArticlesPostRequest(articlesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GolioAPI.ArticlesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArticlesPost`: ArticlesPost200Response
	fmt.Fprintf(os.Stdout, "Response from `GolioAPI.ArticlesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArticlesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **articlesPostRequest** | [**ArticlesPostRequest**](ArticlesPostRequest.md) |  | 

### Return type

[**ArticlesPost200Response**](ArticlesPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnglishDictionaryWordBookmarkDelete

> map[string]interface{} EnglishDictionaryWordBookmarkDelete(ctx, word).Body(body).Execute()

辞書で引いた単語ブックマークを削除



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sunjin110/github.com/sunjin110/folio/lime"
)

func main() {
	word := "word_example" // string | 
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GolioAPI.EnglishDictionaryWordBookmarkDelete(context.Background(), word).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GolioAPI.EnglishDictionaryWordBookmarkDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnglishDictionaryWordBookmarkDelete`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GolioAPI.EnglishDictionaryWordBookmarkDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**word** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnglishDictionaryWordBookmarkDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnglishDictionaryWordBookmarkPut

> map[string]interface{} EnglishDictionaryWordBookmarkPut(ctx, word).Body(body).Execute()

辞書で引いた単語ブックマークを作成



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sunjin110/github.com/sunjin110/folio/lime"
)

func main() {
	word := "word_example" // string | 
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GolioAPI.EnglishDictionaryWordBookmarkPut(context.Background(), word).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GolioAPI.EnglishDictionaryWordBookmarkPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnglishDictionaryWordBookmarkPut`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GolioAPI.EnglishDictionaryWordBookmarkPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**word** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnglishDictionaryWordBookmarkPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnglishDictionaryWordGet

> EnglishDictionaryWordGet200Response EnglishDictionaryWordGet(ctx, word).Execute()

英単語を辞書で引く



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sunjin110/github.com/sunjin110/folio/lime"
)

func main() {
	word := "word_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GolioAPI.EnglishDictionaryWordGet(context.Background(), word).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GolioAPI.EnglishDictionaryWordGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnglishDictionaryWordGet`: EnglishDictionaryWordGet200Response
	fmt.Fprintf(os.Stdout, "Response from `GolioAPI.EnglishDictionaryWordGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**word** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnglishDictionaryWordGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnglishDictionaryWordGet200Response**](EnglishDictionaryWordGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HelloGet

> HelloGet200Response HelloGet(ctx).Execute()

hello



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sunjin110/github.com/sunjin110/folio/lime"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GolioAPI.HelloGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GolioAPI.HelloGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HelloGet`: HelloGet200Response
	fmt.Fprintf(os.Stdout, "Response from `GolioAPI.HelloGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHelloGetRequest struct via the builder pattern


### Return type

[**HelloGet200Response**](HelloGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MediaGet

> MediaGet200Response MediaGet(ctx).Offset(offset).Limit(limit).Execute()

メディア一覧取得



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sunjin110/github.com/sunjin110/folio/lime"
)

func main() {
	offset := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GolioAPI.MediaGet(context.Background()).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GolioAPI.MediaGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MediaGet`: MediaGet200Response
	fmt.Fprintf(os.Stdout, "Response from `GolioAPI.MediaGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**MediaGet200Response**](MediaGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MediaMediumIdDelete

> map[string]interface{} MediaMediumIdDelete(ctx, mediumId).Execute()

メディアの削除



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sunjin110/github.com/sunjin110/folio/lime"
)

func main() {
	mediumId := "mediumId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GolioAPI.MediaMediumIdDelete(context.Background(), mediumId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GolioAPI.MediaMediumIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MediaMediumIdDelete`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GolioAPI.MediaMediumIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mediumId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMediaMediumIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MediaMediumIdGet

> MediaMediumIdGet200Response MediaMediumIdGet(ctx, mediumId).Execute()

メディアの取得



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sunjin110/github.com/sunjin110/folio/lime"
)

func main() {
	mediumId := "mediumId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GolioAPI.MediaMediumIdGet(context.Background(), mediumId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GolioAPI.MediaMediumIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MediaMediumIdGet`: MediaMediumIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `GolioAPI.MediaMediumIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mediumId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMediaMediumIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MediaMediumIdGet200Response**](MediaMediumIdGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MediaPost

> MediaPost200Response MediaPost(ctx).MediaPostRequest(mediaPostRequest).Execute()

メディアの登録



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sunjin110/github.com/sunjin110/folio/lime"
)

func main() {
	mediaPostRequest := *openapiclient.NewMediaPostRequest("FileName_example") // MediaPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GolioAPI.MediaPost(context.Background()).MediaPostRequest(mediaPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GolioAPI.MediaPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MediaPost`: MediaPost200Response
	fmt.Fprintf(os.Stdout, "Response from `GolioAPI.MediaPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mediaPostRequest** | [**MediaPostRequest**](MediaPostRequest.md) |  | 

### Return type

[**MediaPost200Response**](MediaPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksIdGet

> TasksIdGet200Response TasksIdGet(ctx, id).Body(body).Execute()

タスク詳細取得



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sunjin110/github.com/sunjin110/folio/lime"
)

func main() {
	id := "id_example" // string | 
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GolioAPI.TasksIdGet(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GolioAPI.TasksIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksIdGet`: TasksIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `GolioAPI.TasksIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**TasksIdGet200Response**](TasksIdGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksIdPut

> ArticlesPost200Response TasksIdPut(ctx, id).TasksPostRequest(tasksPostRequest).Execute()

タスク編集



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sunjin110/github.com/sunjin110/folio/lime"
)

func main() {
	id := "id_example" // string | 
	tasksPostRequest := *openapiclient.NewTasksPostRequest("Title_example", "Detail_example") // TasksPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GolioAPI.TasksIdPut(context.Background(), id).TasksPostRequest(tasksPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GolioAPI.TasksIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksIdPut`: ArticlesPost200Response
	fmt.Fprintf(os.Stdout, "Response from `GolioAPI.TasksIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tasksPostRequest** | [**TasksPostRequest**](TasksPostRequest.md) |  | 

### Return type

[**ArticlesPost200Response**](ArticlesPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksPost

> ArticlesPost200Response TasksPost(ctx).TasksPostRequest(tasksPostRequest).Execute()

タスク作成



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sunjin110/github.com/sunjin110/folio/lime"
)

func main() {
	tasksPostRequest := *openapiclient.NewTasksPostRequest("Title_example", "Detail_example") // TasksPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GolioAPI.TasksPost(context.Background()).TasksPostRequest(tasksPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GolioAPI.TasksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksPost`: ArticlesPost200Response
	fmt.Fprintf(os.Stdout, "Response from `GolioAPI.TasksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTasksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tasksPostRequest** | [**TasksPostRequest**](TasksPostRequest.md) |  | 

### Return type

[**ArticlesPost200Response**](ArticlesPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranslationPost

> TranslationPost200Response TranslationPost(ctx).TranslationPostRequest(translationPostRequest).Execute()

翻訳



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sunjin110/github.com/sunjin110/folio/lime"
)

func main() {
	translationPostRequest := *openapiclient.NewTranslationPostRequest(openapiclient.language_code("auto"), openapiclient.language_code("auto"), "Text_example") // TranslationPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GolioAPI.TranslationPost(context.Background()).TranslationPostRequest(translationPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GolioAPI.TranslationPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TranslationPost`: TranslationPost200Response
	fmt.Fprintf(os.Stdout, "Response from `GolioAPI.TranslationPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTranslationPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **translationPostRequest** | [**TranslationPostRequest**](TranslationPostRequest.md) |  | 

### Return type

[**TranslationPost200Response**](TranslationPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

