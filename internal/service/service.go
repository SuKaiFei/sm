package service

import (
	"context"
	"errors"

	"github.com/bilibili/kratos/pkg/conf/paladin"
	"sm/internal/dao"
)

const ctxKeyUserInfo = "user-info"

var (
	ErrUserInfoFormat = errors.New("用户信息格式化错误")
)

// Service service.
type Service struct {
	ac  *paladin.Map
	dao *dao.Dao
}

// New new a service and return.
func New() (s *Service) {
	var ac = new(paladin.TOML)
	if err := paladin.Watch("application.toml", ac); err != nil {
		panic(err)
	}
	s = &Service{
		ac:  ac,
		dao: dao.New(),
	}
	return s
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context) (err error) {
	return s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
	s.dao.Close()
}
