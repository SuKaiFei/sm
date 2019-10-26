// 作者 sukaifei
// 日期 2019/09/07

package service

import (
	"context"
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"sm/internal/model"
)

func (s *Service) Register(ctx context.Context, registered *model.Registered) (err error) {
	_, err = s.GetUserByEmail(ctx, registered.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			return
		}
	} else {
		return fmt.Errorf("邮箱(%s)已注册", registered.Email)
	}

	user := new(model.User)
	user.Email = registered.Email
	user.NickName = registered.NickName
	user.Sex = &registered.Sex
	user.Avatar = registered.Avatar
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(registered.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	user.Password = string(fromPassword)
	_, err = s.dao.InsertUser(ctx, user)
	return
}

func (s *Service) GetUserByEmail(ctx context.Context, email string) (user *model.User, err error) {
	return s.dao.GetUserByEmail(ctx, email)
}

func (s *Service) DeleteUserCache(ctx context.Context, authToken string) (err error) {
	return s.dao.DeleteLoginInfo(ctx, authToken)
}

func (s *Service) SaveLoginInfo(ctx context.Context, authToken string, user *model.UserInfo) (err error) {
	return s.dao.SaveLoginInfo(ctx, authToken, user)
}

func (s *Service) GetLoginInfoByToken(ctx context.Context, authToken string) (user *model.UserInfo, err error) {
	return s.dao.GetLoginInfo(ctx, authToken)
}
