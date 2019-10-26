package http

import (
	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"net/http"
	"sm/internal/service"
	"sm/library/net/http/middleware/auth"
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

	var (
		rc struct {
			Conf *redis.Config
		}
	)
	if err := paladin.Get("redis.toml").UnmarshalTOML(&rc); err != nil {
		panic(err)
	}

	auth := auth.New(&auth.Config{Redis: rc.Conf})

	e.Use(cors.Middleware())
	e.Use(pv.PVMiddleware())
	g := e.Group("/sm")
	{
		g.POST("/register", Registered)
		g.Any("/login", Login)
		g.GET("/user", auth.User, UserInfo)
		g.POST("/logout", auth.User, Logout)
	}

	// 文章路由
	g.GET("/articles", GetArticles)
	g.GET("/my/articles", GetMyArticles)
	g.GET("/article/info", GetArticle)
	g.POST("/article/add", auth.User, AddArticle)

	// 标签路由
	g.GET("/tags", GetTags)
	g.POST("/tag/add", AddTag)
	g.DELETE("/tag/delete/:id", DeleteTag)
	g.PUT("/tag/update", UpdateTag)
	g.GET("/tag/get/:id", GetTag)

	file := NewFile()
	g.POST("/upload", file.Upload)
}

func ping(ctx *bm.Context) {
	if err := svc.Ping(ctx); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}
