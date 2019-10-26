// 作者 sukaifei
// 日期 2019/09/10

package cors

import (
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

func Middleware() bm.HandlerFunc {
	return func(c *bm.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("origin"))
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
