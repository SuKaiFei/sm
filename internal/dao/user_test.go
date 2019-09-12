// 作者 sukaifei
// 日期 2019/09/11

package dao

import (
	"context"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestDao_GetUserByEmail(t *testing.T) {
	var (
		c     = context.TODO()
		email = "kaifeisu@gmail.com"
	)

	convey.Convey("GetUserByEmail", t, func(ctx convey.C) {
		_, err := d.GetUserByEmail(c, email)
		ctx.Convey("Then err should be nil.user should not be nil.", func(ctx convey.C) {
			ctx.So(err, convey.ShouldBeNil)
		})
	})
}

func TestDao_GetLoginInfo(t *testing.T) {
	var (
		c     = context.TODO()
		token = "14b085d8-a7a4-4f1b-b987-6ae68c316c6f"
	)

	convey.Convey("GetLoginInfo", t, func(ctx convey.C) {
		userInfo, err := d.GetLoginInfo(c, token)
		ctx.Convey("Then err should be nil.user should not be nil.", func(ctx convey.C) {
			ctx.So(err, convey.ShouldBeNil)
		})
	})
}
