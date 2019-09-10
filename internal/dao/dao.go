package dao

import (
	"context"
	"time"

	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/bilibili/kratos/pkg/log"
	xtime "github.com/bilibili/kratos/pkg/time"
)

// Dao Dao.
type Dao struct {
	db          *sql.DB
	redis       *redis.Pool
	redisExpire int32
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// New new a Dao and return.
func New() *Dao {
	var (
		dc struct {
			Conf *sql.Config
		}
		rc struct {
			Conf       *redis.Config
			ConfExpire xtime.Duration
		}
	)
	checkErr(paladin.Get("mysql.toml").UnmarshalTOML(&dc))
	checkErr(paladin.Get("redis.toml").UnmarshalTOML(&rc))
	return &Dao{
		// mysql
		db: sql.NewMySQL(dc.Conf),
		// redis
		redis:       redis.NewPool(rc.Conf),
		redisExpire: int32(time.Duration(rc.ConfExpire) / time.Second),
	}
}

// Close close the resource.
func (d *Dao) Close() {
	d.redis.Close()
	d.db.Close()
}

// Ping ping the resource.
func (d *Dao) Ping(ctx context.Context) (err error) {
	if err = d.pingRedis(ctx); err != nil {
		return
	}
	return d.db.Ping(ctx)
}

func (d *Dao) pingRedis(ctx context.Context) (err error) {
	conn := d.redis.Get(ctx)
	defer conn.Close()
	if _, err = conn.Do("SET", "ping", "pong"); err != nil {
		log.Error("conn.Set(PING) error(%v)", err)
	}
	return
}
