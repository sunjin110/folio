package model_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/sunjin110/folio/golio/domain/model"
)

func Test_GoogleSearchResult_GetBaseURL(t *testing.T) {
	Convey("Test_GoogleSearchResult_GetBaseURL", t, func() {
		type test struct {
			name    string
			url     string
			want    string
			wantErr error
		}

		tests := []test{
			{
				name: "normal",
				url:  "https://github.com/sunjin110/folio/pull/143",
				want: "https://github.com",
			},
		}

		for _, tt := range tests {
			Convey(tt.name, func() {
				result := &model.GoogleSearchResult{
					URL: tt.url,
				}

				got, err := result.GetBaseURL()
				if tt.wantErr != nil {
					So(err, ShouldBeNil)
					So(err.Error(), ShouldEqual, tt.wantErr.Error())
					return
				}

				So(err, ShouldBeNil)
				So(got, ShouldEqual, tt.want)
			})
		}
	})
}
