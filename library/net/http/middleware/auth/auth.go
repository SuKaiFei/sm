// 作者 sukaifei
// 日期 2019/09/20

package auth

import (
	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/bilibili/kratos/pkg/ecode"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	jsoniter "github.com/json-iterator/go"
	"sm/internal/model"
)

const (
	cookieAuthName = "sm-user"
	ctxKeyUserInfo = "user-info"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Auth struct {
	redis *redis.Pool
}

type Config struct {
	Redis *redis.Config
}

func New(c *Config) (a *Auth) {
	return &Auth{redis: redis.NewPool(c.Redis)}
}

func (a *Auth) User(ctx *bm.Context) {
	req := ctx.Request
	cookie, err := req.Cookie(cookieAuthName)
	if err != nil {
		ctx.JSON(nil, ecode.Unauthorized)
		ctx.Abort()
		return
	}

	conn := a.redis.Get(ctx)
	defer conn.Close()

	userInfoStr, err := redis.String(conn.Do("GET", "auth:"+cookie.Value))
	if err != nil {
		ctx.JSON(nil, ecode.Unauthorized)
		ctx.Abort()
		return
	}
	userInfo := new(model.UserInfo)
	err = json.Unmarshal([]byte(userInfoStr), userInfo)
	if err != nil {
		ctx.JSON(nil, err)
		ctx.Abort()
		return
	}

	ctx.Set(ctxKeyUserInfo, userInfo)
	ctx.Next()
}
