package http

import (
	"net/http"
	"sm/library/net/http/middleware/cors"
	"sm/library/net/http/middleware/pv"

	"sm/internal/service"

	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
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
	engine = bm.DefaultServer(hc.Server)
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
		g.GET("/login", Login)
		g.GET("/tags", GetTags)
		g.POST("/tag/add", AddTag)
		g.DELETE("/tag/delete", DeleteTag)
		g.PUT("/tag/update", UpdateTag)
		g.GET("/tag/get/:id", GetTag)
	}
}

func ping(ctx *bm.Context) {
	if err := svc.Ping(ctx); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}
