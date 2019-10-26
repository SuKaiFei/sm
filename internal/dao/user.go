// 作者 sukaifei
// 日期 2019/09/07

package dao

import (
	"context"
	"fmt"
	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/prometheus/common/log"
	"sm/internal/model"
)

const (
	_selUserByEmail = `SELECT id,avatar,last_login_time,nickname,password,vip,email,heat,sex,role  FROM user where email = ?`
	_insertUser     = "INSERT INTO user(avatar,heat,nickname,role,sex,create_time,email,password,vip,validate_email,last_login_time) VALUES (?,'0',?,'platform',?,UNIX_TIMESTAMP(),?,?,0,1,0)"
)

// 根据邮箱获取对应用户信息
func (d *Dao) InsertUser(ctx context.Context, user *model.User) (id int64, err error) {
	res, err := d.db.Exec(ctx, _insertUser, user.Avatar, user.NickName, user.Sex, user.Email, user.Password)
	if err != nil {
		return
	}
	return res.LastInsertId()
}

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
	key := keyAuth(authToken)
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
	userJsonStr, err := redis.String(conn.Do("GET", "auth:"+authToken))
	if err != nil {
		return
	}
	user = new(model.UserInfo)
	err = json.UnmarshalFromString(userJsonStr, user)
	return
}

// 根据token删除用户信息
func (d *Dao) DeleteLoginInfo(ctx context.Context, authToken string) (err error) {
	conn := d.redis.Get(ctx)
	defer conn.Close()
	_, err = conn.Do("DEL", "auth:"+authToken)
	if err != nil {
		return
	}
	return
}
