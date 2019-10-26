// 作者 sukaifei
// 日期 2019/09/07

package service

import (
	"context"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"sm/internal/model"
	"time"
)

func (s *Service) AddArticle(ctx *bm.Context, req *model.ArticleAdd) (id int64, err error) {
	userO, ok := ctx.Get(ctxKeyUserInfo)
	if !ok {
		return 0, ErrUserInfoFormat
	}
	var (
		userInfo *model.UserInfo
	)
	if userInfo, ok = userO.(*model.UserInfo); !ok {
		return 0, ErrUserInfoFormat
	}
	req.UserId = userInfo.Id
	return s.dao.AddArticle(ctx, req)
}

func (s *Service) GetArticle(ctx context.Context, id int64, textType string) (article *model.ArticleInfo, err error) {
	article, err = s.dao.GetArticle(ctx, id)
	if err != nil {
		return
	}
	article.CreateTime = time.Unix(article.CreateTimeS, 0).Format(model.DefaultTimeFormat)
	if textType == model.MdText {
		article.Text = article.MdText
	} else {
		article.Text = article.HtmlText
	}
	return
}

func (s *Service) GetMyArticles(ctx context.Context, acc string, pn, ps int32) (articles *model.ArticleList, err error) {
	list, err := s.dao.GetMyArticles(ctx, acc, pn, ps)
	if err != nil {
		return
	}
	for _, v := range list.Result {
		v.CreateTime = time.Unix(v.CreateTimeS, 0).Format(model.DefaultTimeFormat)
	}
	return list, err
}

func (s *Service) GetArticles(ctx context.Context, tag string, pn, ps int32) (articles *model.ArticleList, err error) {
	list, err := s.dao.GetArticles(ctx, tag, pn, ps)
	if err != nil {
		return
	}
	for _, v := range list.Result {
		v.CreateTime = time.Unix(v.CreateTimeS, 0).Format(model.DefaultTimeFormat)
	}
	return list, err
}
