// 作者 sukaifei
// 日期 2019/09/07

package http

import (
	"errors"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/bilibili/kratos/pkg/net/http/blademaster/binding"
	"sm/internal/model"
)

var (
	ErrQueryFail = errors.New("查询失败")
)

func AddArticle(bm *bm.Context) {
	add := new(model.ArticleAdd)
	if err := bm.BindWith(add, binding.Default(bm.Request.Method, bm.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	bm.JSON(svc.AddArticle(bm, add))
}

func GetMyArticles(bm *bm.Context) {
	req := new(model.MyArticleListReq)
	if err := bm.BindWith(req, binding.Default(bm.Request.Method, bm.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	bm.JSON(svc.GetMyArticles(bm, req.Acc, req.Pn, req.Ps))
}

func GetArticles(bm *bm.Context) {
	req := new(model.ArticleListReq)
	if err := bm.BindWith(req, binding.Default(bm.Request.Method, bm.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	bm.JSON(svc.GetArticles(bm, req.Tag, req.Pn, req.Ps))
}

func GetArticle(bm *bm.Context) {
	req := new(model.CommonInfoReq)
	if err := bm.BindWith(req, binding.Default(bm.Request.Method, bm.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	textType, _ := bm.Params.Get("text_type")
	bm.JSON(svc.GetArticle(bm, req.Id, textType))
}
