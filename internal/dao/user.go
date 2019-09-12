// 作者 sukaifei
// 日期 2019/09/07

package dao

import (
	"context"
	"fmt"
	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/prometheus/common/log"
	"sm/internal/model"
)

const (
	_selUserByEmail = `SELECT id,avatar,last_login_time,nickname,password,vip,email,heat,sex,role  FROM user where email = ?`
)

// 根据邮箱获取对应用户信息
func (d *Dao) GetUserByEmail(ctx context.Context, email string) (user *model.User, err error) {
	row := d.db.QueryRow(ctx, _selUserByEmail, email)
	user = new(model.User)
	err = row.Scan(&user.Id, &user.Avatar, &user.LastLoginTime, &user.NickName,
		&user.Password, &user.Vip, &user.Email, &user.Heat, &user.Sex, &user.Role)
	if err != nil {
		return
	}
	return
}

// 根据token存入用户信息
func (d *Dao) SaveLoginInfo(ctx context.Context, authToken string, user *model.UserInfo) (err error) {
	conn := d.redis.Get(ctx)
	defer conn.Close()
	key := keyAuth(user.Email, authToken)
	userJson, err := json.Marshal(user)
	if err != nil {
		return
	}
	_, err = conn.Do("SET", key, userJson, "EX", _redisAuthExpire)
	if err != nil {
		log.Error(fmt.Sprintf("conn.Set(%s) error(%v)", key, err))
	}
	return
}

// 根据token获取用户信息
func (d *Dao) GetLoginInfo(ctx context.Context, authToken string) (user *model.UserInfo, err error) {
	conn := d.redis.Get(ctx)
	defer conn.Close()
	key := "*:*:" + authToken

	keys, err := redis.Strings(conn.Do("KEYS", key))
	if err != nil {
		log.Error(fmt.Sprintf("conn.Get(%s) error(%v)", key, err))
	}
	if len(keys) == 0 {
		err = sql.ErrNoRows
		return
	}
	key = keys[0]
	userJsonStr, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return
	}
	user = new(model.UserInfo)
	err = json.UnmarshalFromString(userJsonStr, user)
	return
}
