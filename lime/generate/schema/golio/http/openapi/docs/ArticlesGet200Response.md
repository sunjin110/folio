# ArticlesGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Articles** | [**[]Article**](Article.md) |  | 
**Total** | **int32** |  | 

## Methods

### NewArticlesGet200Response

`func NewArticlesGet200Response(articles []Article, total int32, ) *ArticlesGet200Response`

NewArticlesGet200Response instantiates a new ArticlesGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArticlesGet200ResponseWithDefaults

`func NewArticlesGet200ResponseWithDefaults() *ArticlesGet200Response`

NewArticlesGet200ResponseWithDefaults instantiates a new ArticlesGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArticles

`func (o *ArticlesGet200Response) GetArticles() []Article`

GetArticles returns the Articles field if non-nil, zero value otherwise.

### GetArticlesOk

`func (o *ArticlesGet200Response) GetArticlesOk() (*[]Article, bool)`

GetArticlesOk returns a tuple with the Articles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticles

`func (o *ArticlesGet200Response) SetArticles(v []Article)`

SetArticles sets Articles field to given value.


### GetTotal

`func (o *ArticlesGet200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ArticlesGet200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ArticlesGet200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


