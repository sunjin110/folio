// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * folio
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

package openapi


import (
	"time"
)



type TasksPostRequest struct {

	Title string `json:"title"`

	Detail string `json:"detail"`

	StartTime *time.Time `json:"start_time,omitempty"`

	DueTime *time.Time `json:"due_time,omitempty"`
}

// AssertTasksPostRequestRequired checks if the required fields are not zero-ed
func AssertTasksPostRequestRequired(obj TasksPostRequest) error {
	elements := map[string]interface{}{
		"title": obj.Title,
		"detail": obj.Detail,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertTasksPostRequestConstraints checks if the values respects the defined constraints
func AssertTasksPostRequestConstraints(obj TasksPostRequest) error {
	return nil
}