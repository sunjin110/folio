# TasksIdGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Title** | **string** |  | 
**Detail** | **string** |  | 
**StartTime** | Pointer to **NullableTime** |  | [optional] 
**DueTime** | Pointer to **NullableTime** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewTasksIdGet200Response

`func NewTasksIdGet200Response(id string, title string, detail string, createdAt time.Time, updatedAt time.Time, ) *TasksIdGet200Response`

NewTasksIdGet200Response instantiates a new TasksIdGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTasksIdGet200ResponseWithDefaults

`func NewTasksIdGet200ResponseWithDefaults() *TasksIdGet200Response`

NewTasksIdGet200ResponseWithDefaults instantiates a new TasksIdGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TasksIdGet200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TasksIdGet200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TasksIdGet200Response) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *TasksIdGet200Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TasksIdGet200Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TasksIdGet200Response) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDetail

`func (o *TasksIdGet200Response) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *TasksIdGet200Response) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *TasksIdGet200Response) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetStartTime

`func (o *TasksIdGet200Response) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TasksIdGet200Response) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TasksIdGet200Response) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TasksIdGet200Response) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *TasksIdGet200Response) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *TasksIdGet200Response) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetDueTime

`func (o *TasksIdGet200Response) GetDueTime() time.Time`

GetDueTime returns the DueTime field if non-nil, zero value otherwise.

### GetDueTimeOk

`func (o *TasksIdGet200Response) GetDueTimeOk() (*time.Time, bool)`

GetDueTimeOk returns a tuple with the DueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueTime

`func (o *TasksIdGet200Response) SetDueTime(v time.Time)`

SetDueTime sets DueTime field to given value.

### HasDueTime

`func (o *TasksIdGet200Response) HasDueTime() bool`

HasDueTime returns a boolean if a field has been set.

### SetDueTimeNil

`func (o *TasksIdGet200Response) SetDueTimeNil(b bool)`

 SetDueTimeNil sets the value for DueTime to be an explicit nil

### UnsetDueTime
`func (o *TasksIdGet200Response) UnsetDueTime()`

UnsetDueTime ensures that no value is present for DueTime, not even an explicit nil
### GetCreatedAt

`func (o *TasksIdGet200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TasksIdGet200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TasksIdGet200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *TasksIdGet200Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TasksIdGet200Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TasksIdGet200Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


