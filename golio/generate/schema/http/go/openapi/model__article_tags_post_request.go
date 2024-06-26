// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * folio
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

package openapi




type ArticleTagsPostRequest struct {

	Name string `json:"name"`
}

// AssertArticleTagsPostRequestRequired checks if the required fields are not zero-ed
func AssertArticleTagsPostRequestRequired(obj ArticleTagsPostRequest) error {
	elements := map[string]interface{}{
		"name": obj.Name,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertArticleTagsPostRequestConstraints checks if the values respects the defined constraints
func AssertArticleTagsPostRequestConstraints(obj ArticleTagsPostRequest) error {
	return nil
}
