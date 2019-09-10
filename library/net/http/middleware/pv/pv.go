// 作者 sukaifei
// 日期 2019/09/10

package pv

import (
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

func PVMiddleware() bm.HandlerFunc {
	return func(bm *bm.Context) {
		if bm.Request.Method != "OPTIONS" {

		}
	}
}
