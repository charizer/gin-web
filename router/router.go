package router

import (
	"github.com/chenzr/gin-web/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

// router.go 所有的路由放到这里

func Load(g *gin.Engine) *gin.Engine {
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})
	group := g.Group("/api/v1/web")
	{
		group.GET("/req/:param", handler.Handler)
	}
	return g
}
