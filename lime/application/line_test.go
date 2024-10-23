package application_test

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/sunjin110/folio/lime/application"
)

func Test_lineUsecase_VerifySignature(t *testing.T) {
	SkipConvey("Test_lineUsecase_VerifySignature", t, func() {

		lineUsecase := application.NewLineUsecase("line channel secret")
		err := lineUsecase.VerifySignature(context.Background(), "signature", []byte("{\"destination\":\"Uc991ef780f1804f2323b8ce0c7ec65ba\",\"events\":[]}"))
		So(err, ShouldBeNil)
	})
}
