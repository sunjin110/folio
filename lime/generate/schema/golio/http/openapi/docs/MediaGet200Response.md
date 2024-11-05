# MediaGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** |  | 
**Media** | [**[]MediaGet200ResponseMediaInner**](MediaGet200ResponseMediaInner.md) |  | 

## Methods

### NewMediaGet200Response

`func NewMediaGet200Response(totalCount int32, media []MediaGet200ResponseMediaInner, ) *MediaGet200Response`

NewMediaGet200Response instantiates a new MediaGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaGet200ResponseWithDefaults

`func NewMediaGet200ResponseWithDefaults() *MediaGet200Response`

NewMediaGet200ResponseWithDefaults instantiates a new MediaGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *MediaGet200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *MediaGet200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *MediaGet200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetMedia

`func (o *MediaGet200Response) GetMedia() []MediaGet200ResponseMediaInner`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *MediaGet200Response) GetMediaOk() (*[]MediaGet200ResponseMediaInner, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *MediaGet200Response) SetMedia(v []MediaGet200ResponseMediaInner)`

SetMedia sets Media field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


