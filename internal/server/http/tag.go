// 作者 sukaifei
// 日期 2019/09/07

package http

import (
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

func GetTags(bm *bm.Context) {
	tags, err := svc.GetTags(bm)
	bm.JSON(tags, err)
}
