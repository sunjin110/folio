// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * folio
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

package openapi




type WordDetail struct {

	Word string `json:"word"`

	Definitions []WordDefinition `json:"definitions"`

	Frequency float32 `json:"frequency"`
}

// AssertWordDetailRequired checks if the required fields are not zero-ed
func AssertWordDetailRequired(obj WordDetail) error {
	elements := map[string]interface{}{
		"word": obj.Word,
		"definitions": obj.Definitions,
		"frequency": obj.Frequency,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Definitions {
		if err := AssertWordDefinitionRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertWordDetailConstraints checks if the values respects the defined constraints
func AssertWordDetailConstraints(obj WordDetail) error {
	for _, el := range obj.Definitions {
		if err := AssertWordDefinitionConstraints(el); err != nil {
			return err
		}
	}
	return nil
}