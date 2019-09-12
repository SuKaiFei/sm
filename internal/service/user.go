// 作者 sukaifei
// 日期 2019/09/07

package service

import (
	"context"
	"sm/internal/model"
)

func (s *Service) GetUserByEmail(ctx context.Context, email string) (user *model.User, err error) {
	return s.dao.GetUserByEmail(ctx, email)
}

func (s *Service) SaveLoginInfo(ctx context.Context, authToken string, user *model.UserInfo) (err error) {
	return s.dao.SaveLoginInfo(ctx, authToken, user)
}

func (s *Service) GetLoginInfoByToken(ctx context.Context, authToken string) (user *model.UserInfo, err error) {
	return s.dao.GetLoginInfo(ctx, authToken)
}
