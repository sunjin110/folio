package model_test

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/sunjin110/folio/golio/domain/model"
)

// go test -v -count=1 -timeout 30s -run ^TestHtmlExtracter_ExtractText$ github.com/sunjin110/folio/golio/domain/model
func TestHtmlExtractor_ExtractText(t *testing.T) {
	Convey("TestHtmlExtracter_ExtractText", t, func() {
		type test struct {
			name        string
			htmlContent string
			want        string
			wantErr     error
		}

		tests := []test{
			{
				name:        "取得できていること",
				htmlContent: "<html><head><title>Test</title></head><body><p>This is a test paragraph.</p></body></html>",
				want:        "Test\nThis is a test paragraph.\n",
			},
		}

		for _, tt := range tests {
			Convey(tt.name, func() {
				extracter := model.NewHtmlExtractor(tt.htmlContent)

				got, err := extracter.ExtractText(context.Background())
				if tt.wantErr != nil {
					So(err, ShouldBeError)
					So(err.Error(), ShouldEqual, tt.wantErr.Error())
					return
				}

				So(err, ShouldBeNil)
				So(got, ShouldEqual, tt.want)
			})
		}
	})
}
