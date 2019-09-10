// 作者 sukaifei
// 日期 2019/09/10

package netutil

import (
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"strings"
)

func GetClientIP(ctx *bm.Context) string {
	ip := ctx.Request.Header.Get("X-Forwarded-For")

	if strings.Contains(ip, "127.0.0.1") || ip == "" {
		ip = ctx.Request.Header.Get("X-real-ip")
	}
	if ip == "" {
		return "127.0.0.1"
	}

	return ip
}
