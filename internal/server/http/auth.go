// 作者 sukaifei
// 日期 2019/09/07

package http

import (
	"errors"
	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/bilibili/kratos/pkg/ecode"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/bilibili/kratos/pkg/net/http/blademaster/binding"
	"github.com/bilibili/kratos/pkg/net/http/blademaster/render"
	"golang.org/x/crypto/bcrypt"
	"sm/internal/model"
)

var ErrLoginFail = errors.New("用户名和密码错误")
var ErrRegisteredAccountAlreadyExists = errors.New("邮箱已存在")

const (
	cookieAuthName   = "sm-user"
	cookieAuthMaxAge = 60 * 60 * 24 * 30
)

// 注册
func Registered(bm *bm.Context) {

}

// 登录
func Login(bm *bm.Context) {
	login := new(model.Login)
	if err := bm.BindWith(login, binding.Default(bm.Request.Method, bm.Request.Header.Get("Content-Type"))); err != nil {
		bm.JSON(nil, err)
		return
	}

	user, err := svc.GetUserByEmail(bm, login.UserName)
	if err != nil {
		if err == sql.ErrNoRows {
			bm.Render(ecode.RequestErr.Code(), render.JSON{
				Code:    ecode.RequestErr.Code(),
				Message: ErrLoginFail.Error(),
				Data:    nil,
			})
		} else {
			bm.JSON(nil, err)
		}
		return
	}

	isPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))

	if isPassword != nil {
		bm.Render(ecode.RequestErr.Code(), render.JSON{
			Code:    ecode.RequestErr.Code(),
			Message: ErrLoginFail.Error(),
			Data:    nil,
		})
		return
	}

	userInfo := &model.UserInfo{
		Id:            user.Id,
		NickName:      user.NickName,
		LastLoginTime: user.LastLoginTime,
		Role:          user.Role,
		Vip:           user.Vip,
		Email:         user.Email,
		Sex:           user.Sex,
		Heat:          user.Heat,
		Avatar:        user.Avatar,
	}

	authToken, err := bm.Request.Cookie(cookieAuthName)
	if err == nil {
		err = svc.SaveLoginInfo(bm, authToken.Value, userInfo)
		if err != nil {
			bm.JSON(userInfo, err)
			return
		}
	}

	bm.JSON(userInfo, nil)
}

func UserInfo(bm *bm.Context) {
	cookie, err := bm.Request.Cookie(cookieAuthName)
	if err != nil {
		bm.JSON(nil, err)
		return
	}
	bm.JSON(svc.GetLoginInfoByToken(bm, cookie.Value))
}
