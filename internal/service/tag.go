// 作者 sukaifei
// 日期 2019/09/07

package service

import (
	"context"
	"sm/internal/model"
)

func (s *Service) GetTags(ctx context.Context) (tags []*model.Tags, err error) {
	select {
	default:
	case <-ctx.Done():
		return nil, ctx.Err()
	}
	return s.dao.GetTags(ctx)
}
