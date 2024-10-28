package model_test

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/sunjin110/folio/gomb/domain/model"
)

func Test_Image(t *testing.T) {
	Convey("Test_Image", t, func() {
		f, err := os.Open("testdata/test.png")
		So(err, ShouldBeNil)
		defer f.Close()

		img, err := model.NewImage(f)
		So(err, ShouldBeNil)
		thumbnail := img.Thumbnail()

		thumbF, err := os.Create("testdata/test_thumbnail.jpeg")
		So(err, ShouldBeNil)

		err = thumbnail.EncodeJpeg(thumbF)
		So(err, ShouldBeNil)
	})

}
