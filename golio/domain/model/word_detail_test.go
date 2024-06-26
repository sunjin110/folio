package model_test

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/sunjin110/folio/golio/domain/model"
)

// go test -v -count=1 -timeout 30s -run ^Test_WordDetail_ToMarkdown$ github.com/sunjin110/folio/golio/domain/model
func Test_WordDetail_ToMarkdown(t *testing.T) {
	Convey("Test_WordDetail_ToMarkdown", t, func() {
		wordDetail := &model.WordDetail{
			Word: "pen",
			Definitions: []*model.WordDefinition{
				{
					Definition:   "definition_1",
					PartOfSpeech: "noun",
					Synonyms: []string{
						"synonyms_1",
						"synonyms_2",
						"synonyms_3",
					},
					Antonyms: []string{
						"antonyms_1",
						"antonyms_2",
						"antonyms_3",
					},
					Examples: []string{
						"example_1",
						"example_2",
						"example_3",
					},
				},
				{
					Definition:   "definition_2",
					PartOfSpeech: "verb",
					Synonyms:     []string{},
					Antonyms:     []string{},
					Examples:     []string{},
				},
			},
			Frequency: 10.0,
		}

		result, err := wordDetail.ToMarkdown()
		So(err, ShouldBeNil)
		fmt.Println("result is ", result)
	})
}
