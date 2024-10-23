package presentation_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/sunjin110/folio/lime/presentation"
)

func Test_GetHandler(t *testing.T) {
	Convey("Test_GetHandler", t, func() {
		Convey("panicしないこと", func() {
			_, err := presentation.GetHandler()
			So(err, ShouldBeNil)
		})
	})
}
