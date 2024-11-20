# WordDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Word** | **string** |  | 
**Definitions** | [**[]WordDefinition**](WordDefinition.md) |  | 
**Frequency** | **float32** |  | 

## Methods

### NewWordDetail

`func NewWordDetail(word string, definitions []WordDefinition, frequency float32, ) *WordDetail`

NewWordDetail instantiates a new WordDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWordDetailWithDefaults

`func NewWordDetailWithDefaults() *WordDetail`

NewWordDetailWithDefaults instantiates a new WordDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWord

`func (o *WordDetail) GetWord() string`

GetWord returns the Word field if non-nil, zero value otherwise.

### GetWordOk

`func (o *WordDetail) GetWordOk() (*string, bool)`

GetWordOk returns a tuple with the Word field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWord

`func (o *WordDetail) SetWord(v string)`

SetWord sets Word field to given value.


### GetDefinitions

`func (o *WordDetail) GetDefinitions() []WordDefinition`

GetDefinitions returns the Definitions field if non-nil, zero value otherwise.

### GetDefinitionsOk

`func (o *WordDetail) GetDefinitionsOk() (*[]WordDefinition, bool)`

GetDefinitionsOk returns a tuple with the Definitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitions

`func (o *WordDetail) SetDefinitions(v []WordDefinition)`

SetDefinitions sets Definitions field to given value.


### GetFrequency

`func (o *WordDetail) GetFrequency() float32`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *WordDetail) GetFrequencyOk() (*float32, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *WordDetail) SetFrequency(v float32)`

SetFrequency sets Frequency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


