package model_test

import (
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/sunjin110/folio/lime/domain/model"
)

// go test -v -count=1 -timeout 30s -run ^Test_Content_FileName$ github.com/sunjin110/folio/lime/domain/model
func Test_Content_FileName(t *testing.T) {
	Convey("Test_Content_FileName", t, func() {
		type args struct {
			mediaType string
			fileName  *string
		}

		type test struct {
			name    string
			args    args
			wantExt string
			wantErr error
		}

		tests := []test{
			{
				name: "html",
				args: args{
					mediaType: "text/html",
				},
				wantExt: "htm",
			},
			{
				name: "jpg",
				args: args{
					mediaType: "image/jpeg",
				},
				wantExt: "jpe",
			},
		}

		for _, tt := range tests {
			Convey(tt.name, func() {

				c := model.NewContent(nil, tt.args.mediaType, 0, tt.args.fileName)

				got, err := c.FileName()
				if tt.wantErr != nil {
					So(err, ShouldBeError)
					So(err.Error(), ShouldEqual, tt.wantErr.Error())
					return
				}
				So(err, ShouldBeNil)
				s := strings.Split(got, ".")
				So(s[len(s)-1], ShouldEqual, tt.wantExt)
			})
		}
	})
}
