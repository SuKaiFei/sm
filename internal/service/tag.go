// 作者 sukaifei
// 日期 2019/09/07

package service

import (
	"context"
	"sm/internal/model"
)

func (s *Service) GetTags(ctx context.Context) (tags []*model.Tags, err error) {
	return s.dao.GetTags(ctx)
}

func (s *Service) AddTag(ctx context.Context, res *model.Tag) (id int64, err error) {
	return s.dao.AddTag(ctx, res)
}

func (s *Service) DeleteTag(ctx context.Context, id int64) (err error) {
	return s.dao.DeleteTag(ctx, id)
}

func (s *Service) UpdateTag(ctx context.Context, res *model.Tag) (id int64, err error) {
	return s.dao.UpdateTag(ctx, res)
}

func (s *Service) GetTag(ctx context.Context, id int64) (tags *model.Tag, err error) {
	return s.dao.GetTag(ctx, id)
}
