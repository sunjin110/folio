# WordDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Definition** | **string** |  | 
**PartOfSpeech** | **string** |  | 
**Synonyms** | **[]string** |  | 
**Antonyms** | **[]string** |  | 
**Examples** | **[]string** |  | 

## Methods

### NewWordDefinition

`func NewWordDefinition(definition string, partOfSpeech string, synonyms []string, antonyms []string, examples []string, ) *WordDefinition`

NewWordDefinition instantiates a new WordDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWordDefinitionWithDefaults

`func NewWordDefinitionWithDefaults() *WordDefinition`

NewWordDefinitionWithDefaults instantiates a new WordDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefinition

`func (o *WordDefinition) GetDefinition() string`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *WordDefinition) GetDefinitionOk() (*string, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *WordDefinition) SetDefinition(v string)`

SetDefinition sets Definition field to given value.


### GetPartOfSpeech

`func (o *WordDefinition) GetPartOfSpeech() string`

GetPartOfSpeech returns the PartOfSpeech field if non-nil, zero value otherwise.

### GetPartOfSpeechOk

`func (o *WordDefinition) GetPartOfSpeechOk() (*string, bool)`

GetPartOfSpeechOk returns a tuple with the PartOfSpeech field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartOfSpeech

`func (o *WordDefinition) SetPartOfSpeech(v string)`

SetPartOfSpeech sets PartOfSpeech field to given value.


### GetSynonyms

`func (o *WordDefinition) GetSynonyms() []string`

GetSynonyms returns the Synonyms field if non-nil, zero value otherwise.

### GetSynonymsOk

`func (o *WordDefinition) GetSynonymsOk() (*[]string, bool)`

GetSynonymsOk returns a tuple with the Synonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonyms

`func (o *WordDefinition) SetSynonyms(v []string)`

SetSynonyms sets Synonyms field to given value.


### GetAntonyms

`func (o *WordDefinition) GetAntonyms() []string`

GetAntonyms returns the Antonyms field if non-nil, zero value otherwise.

### GetAntonymsOk

`func (o *WordDefinition) GetAntonymsOk() (*[]string, bool)`

GetAntonymsOk returns a tuple with the Antonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntonyms

`func (o *WordDefinition) SetAntonyms(v []string)`

SetAntonyms sets Antonyms field to given value.


### GetExamples

`func (o *WordDefinition) GetExamples() []string`

GetExamples returns the Examples field if non-nil, zero value otherwise.

### GetExamplesOk

`func (o *WordDefinition) GetExamplesOk() (*[]string, bool)`

GetExamplesOk returns a tuple with the Examples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExamples

`func (o *WordDefinition) SetExamples(v []string)`

SetExamples sets Examples field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


