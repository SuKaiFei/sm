package http

import (
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"net/http"
	"sm/internal/service"
	"sm/library/net/http/middleware/cors"
	"sm/library/net/http/middleware/pv"
)

var (
	svc *service.Service
)

// New new a bm server.
func New(s *service.Service) (engine *bm.Engine) {
	var (
		hc struct {
			Server *bm.ServerConfig
		}
	)
	if err := paladin.Get("http.toml").UnmarshalTOML(&hc); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}
	svc = s
	engine = bm.NewServer(hc.Server)
	engine.Use(bm.Recovery(), bm.Trace(), bm.Logger())
	// 挂载自适应限流中间件到 bm engine，使用默认配置
	limiter := bm.NewRateLimiter(nil)
	engine.Use(limiter.Limit())
	initRouter(engine)
	if err := engine.Start(); err != nil {
		panic(err)
	}
	return
}

func initRouter(e *bm.Engine) {
	e.Ping(ping)

	e.Use(cors.CORSMiddleware())
	e.Use(pv.PVMiddleware())
	g := e.Group("/sm")
	{
		g.Any("/login", Login)
		g.GET("/user", UserInfo)
		g.GET("/tags", GetTags)
	}

	file := NewFile()
	g.POST("/upload", file.Upload)
}

func ping(ctx *bm.Context) {
	if err := svc.Ping(ctx); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}
