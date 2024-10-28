package model_test

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/sunjin110/folio/gomb/domain/model"
)

func TestVideo(t *testing.T) {
	Convey("TestVideo", t, func() {

		videoFile, err := os.Open("testdata/test.mp4")
		So(err, ShouldBeNil)

		video, err := model.NewVideo(videoFile)
		So(err, ShouldBeNil)
		thumbnail, err := video.Thumbnail()
		So(err, ShouldBeNil)

		videoThumbnailFile, err := os.Create("testdata/test_video_thumbnail.jpeg")
		So(err, ShouldBeNil)

		err = thumbnail.EncodeJpeg(videoThumbnailFile)
		So(err, ShouldBeNil)
	})
}
