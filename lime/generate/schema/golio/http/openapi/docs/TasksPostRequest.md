# TasksPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Detail** | **string** |  | 
**StartTime** | Pointer to **NullableTime** |  | [optional] 
**DueTime** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewTasksPostRequest

`func NewTasksPostRequest(title string, detail string, ) *TasksPostRequest`

NewTasksPostRequest instantiates a new TasksPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTasksPostRequestWithDefaults

`func NewTasksPostRequestWithDefaults() *TasksPostRequest`

NewTasksPostRequestWithDefaults instantiates a new TasksPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *TasksPostRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TasksPostRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TasksPostRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDetail

`func (o *TasksPostRequest) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *TasksPostRequest) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *TasksPostRequest) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetStartTime

`func (o *TasksPostRequest) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TasksPostRequest) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TasksPostRequest) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TasksPostRequest) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *TasksPostRequest) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *TasksPostRequest) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetDueTime

`func (o *TasksPostRequest) GetDueTime() time.Time`

GetDueTime returns the DueTime field if non-nil, zero value otherwise.

### GetDueTimeOk

`func (o *TasksPostRequest) GetDueTimeOk() (*time.Time, bool)`

GetDueTimeOk returns a tuple with the DueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueTime

`func (o *TasksPostRequest) SetDueTime(v time.Time)`

SetDueTime sets DueTime field to given value.

### HasDueTime

`func (o *TasksPostRequest) HasDueTime() bool`

HasDueTime returns a boolean if a field has been set.

### SetDueTimeNil

`func (o *TasksPostRequest) SetDueTimeNil(b bool)`

 SetDueTimeNil sets the value for DueTime to be an explicit nil

### UnsetDueTime
`func (o *TasksPostRequest) UnsetDueTime()`

UnsetDueTime ensures that no value is present for DueTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


