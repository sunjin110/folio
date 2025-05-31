# ArticlesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Body** | **string** |  | 
**TagIds** | **[]string** |  | 

## Methods

### NewArticlesPostRequest

`func NewArticlesPostRequest(title string, body string, tagIds []string, ) *ArticlesPostRequest`

NewArticlesPostRequest instantiates a new ArticlesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArticlesPostRequestWithDefaults

`func NewArticlesPostRequestWithDefaults() *ArticlesPostRequest`

NewArticlesPostRequestWithDefaults instantiates a new ArticlesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ArticlesPostRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ArticlesPostRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ArticlesPostRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBody

`func (o *ArticlesPostRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ArticlesPostRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ArticlesPostRequest) SetBody(v string)`

SetBody sets Body field to given value.


### GetTagIds

`func (o *ArticlesPostRequest) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *ArticlesPostRequest) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *ArticlesPostRequest) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


