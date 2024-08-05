package repository_test

import (
	"context"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/sunjin110/folio/golio/infrastructure/repository"
)

func Test_oauth_GenerateAuthorizationURL(t *testing.T) {
	Convey("Test_oauth_GenerateAuthorizationURL", t, func() {

	})
}

// go test -v -count=1 -timeout 30s -run ^Test_googleOauth2_VerifyToken_Real$ github.com/sunjin110/folio/golio/infrastructure/repository
func Test_googleOauth2_VerifyToken_Real(t *testing.T) {
	SkipConvey("Test_googleOauth2_VerifyToken_Real", t, func() {
		repo, err := repository.NewGoogleOAuth2(context.Background(), "", "", "")
		So(err, ShouldBeNil)

		token := "xxx"

		ok, expireTime, err := repo.VerifyToken(context.Background(), token)
		So(err, ShouldBeNil)
		So(ok, ShouldBeTrue)
		fmt.Println("ok is ", ok)
		fmt.Println("expireTime is ", expireTime)
	})
}
