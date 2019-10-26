// 作者 sukaifei
// 日期 2019/09/11

package dao

import "fmt"

const (
	_prefixAuth = "auth:%s"

	_redisAuthExpire = 60 * 60 * 24 * 30
)

func keyAuth(token string) string {
	return fmt.Sprintf(_prefixAuth, token)
}
