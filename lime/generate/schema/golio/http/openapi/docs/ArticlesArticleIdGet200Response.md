# ArticlesArticleIdGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Title** | **string** |  | 
**Body** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UserId** | **string** |  | 
**Tags** | [**[]ArticleTag**](ArticleTag.md) |  | 

## Methods

### NewArticlesArticleIdGet200Response

`func NewArticlesArticleIdGet200Response(id string, title string, body string, createdAt time.Time, userId string, tags []ArticleTag, ) *ArticlesArticleIdGet200Response`

NewArticlesArticleIdGet200Response instantiates a new ArticlesArticleIdGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArticlesArticleIdGet200ResponseWithDefaults

`func NewArticlesArticleIdGet200ResponseWithDefaults() *ArticlesArticleIdGet200Response`

NewArticlesArticleIdGet200ResponseWithDefaults instantiates a new ArticlesArticleIdGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ArticlesArticleIdGet200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArticlesArticleIdGet200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArticlesArticleIdGet200Response) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *ArticlesArticleIdGet200Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ArticlesArticleIdGet200Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ArticlesArticleIdGet200Response) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBody

`func (o *ArticlesArticleIdGet200Response) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ArticlesArticleIdGet200Response) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ArticlesArticleIdGet200Response) SetBody(v string)`

SetBody sets Body field to given value.


### GetCreatedAt

`func (o *ArticlesArticleIdGet200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArticlesArticleIdGet200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArticlesArticleIdGet200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUserId

`func (o *ArticlesArticleIdGet200Response) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ArticlesArticleIdGet200Response) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ArticlesArticleIdGet200Response) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetTags

`func (o *ArticlesArticleIdGet200Response) GetTags() []ArticleTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ArticlesArticleIdGet200Response) GetTagsOk() (*[]ArticleTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ArticlesArticleIdGet200Response) SetTags(v []ArticleTag)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


