// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * folio
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

package openapi




type EnglishDictionaryWordGet200Response struct {

	Origin WordDetail `json:"origin"`

	Translated WordDetail `json:"translated"`
}

// AssertEnglishDictionaryWordGet200ResponseRequired checks if the required fields are not zero-ed
func AssertEnglishDictionaryWordGet200ResponseRequired(obj EnglishDictionaryWordGet200Response) error {
	elements := map[string]interface{}{
		"origin": obj.Origin,
		"translated": obj.Translated,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertWordDetailRequired(obj.Origin); err != nil {
		return err
	}
	if err := AssertWordDetailRequired(obj.Translated); err != nil {
		return err
	}
	return nil
}

// AssertEnglishDictionaryWordGet200ResponseConstraints checks if the values respects the defined constraints
func AssertEnglishDictionaryWordGet200ResponseConstraints(obj EnglishDictionaryWordGet200Response) error {
	if err := AssertWordDetailConstraints(obj.Origin); err != nil {
		return err
	}
	if err := AssertWordDetailConstraints(obj.Translated); err != nil {
		return err
	}
	return nil
}
