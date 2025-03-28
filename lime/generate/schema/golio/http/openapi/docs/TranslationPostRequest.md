# TranslationPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceLanguageCode** | [**LanguageCode**](LanguageCode.md) |  | 
**TargetLanguageCode** | [**LanguageCode**](LanguageCode.md) |  | 
**Text** | **string** | 翻訳するテキスト | 

## Methods

### NewTranslationPostRequest

`func NewTranslationPostRequest(sourceLanguageCode LanguageCode, targetLanguageCode LanguageCode, text string, ) *TranslationPostRequest`

NewTranslationPostRequest instantiates a new TranslationPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranslationPostRequestWithDefaults

`func NewTranslationPostRequestWithDefaults() *TranslationPostRequest`

NewTranslationPostRequestWithDefaults instantiates a new TranslationPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceLanguageCode

`func (o *TranslationPostRequest) GetSourceLanguageCode() LanguageCode`

GetSourceLanguageCode returns the SourceLanguageCode field if non-nil, zero value otherwise.

### GetSourceLanguageCodeOk

`func (o *TranslationPostRequest) GetSourceLanguageCodeOk() (*LanguageCode, bool)`

GetSourceLanguageCodeOk returns a tuple with the SourceLanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLanguageCode

`func (o *TranslationPostRequest) SetSourceLanguageCode(v LanguageCode)`

SetSourceLanguageCode sets SourceLanguageCode field to given value.


### GetTargetLanguageCode

`func (o *TranslationPostRequest) GetTargetLanguageCode() LanguageCode`

GetTargetLanguageCode returns the TargetLanguageCode field if non-nil, zero value otherwise.

### GetTargetLanguageCodeOk

`func (o *TranslationPostRequest) GetTargetLanguageCodeOk() (*LanguageCode, bool)`

GetTargetLanguageCodeOk returns a tuple with the TargetLanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLanguageCode

`func (o *TranslationPostRequest) SetTargetLanguageCode(v LanguageCode)`

SetTargetLanguageCode sets TargetLanguageCode field to given value.


### GetText

`func (o *TranslationPostRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TranslationPostRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TranslationPostRequest) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


