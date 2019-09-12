// 作者 sukaifei
// 日期 2019/09/07

package http

import (
	"errors"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/bilibili/kratos/pkg/net/http/blademaster/binding"
	"sm/internal/model"
	"strconv"
)

func GetTags(bm *bm.Context) {
	tags, err := svc.GetTags(bm)
	bm.JSON(tags, err)
}

func AddTag(bm *bm.Context) {
	tag := new(model.Tag)
	if err := bm.BindWith(tag, binding.Default(bm.Request.Method, bm.Request.Header.Get("Content-Type"))); err != nil {
		bm.JSON(nil, err)
		return
	}
	id, err := svc.AddTag(bm, tag)
	bm.JSON(id, err)
}

func DeleteTag(bm *bm.Context) {
	id, exists := bm.Get("id")
	if !exists {
		bm.JSON(nil, errors.New("删除失败"))
		return
	}
	err := svc.DeleteTag(bm, id.(*int64))
	bm.JSON(id, err)
}

func UpdateTag(bm *bm.Context) {
	tag := new(model.Tag)
	if err := bm.BindWith(tag, binding.Default(bm.Request.Method, bm.Request.Header.Get("Content-Type"))); err != nil {
		bm.JSON(nil, err)
		return
	}
	id, err := svc.UpdateTag(bm, tag)
	bm.JSON(id, err)
}

func GetTag(bm *bm.Context) {
	params, exists := bm.Params.Get("id")
	if !exists {
		bm.JSON(nil, errors.New("查询失败"))
		return
	}
	id, err := strconv.ParseInt(params, 10, 64)

	tag, err := svc.GetTag(bm, &id)
	bm.JSON(tag, err)
}
