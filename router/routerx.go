package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kiririx/krim/api"
	"github.com/kiririx/krim/module/req"
	"github.com/kiririx/krim/util/callback"
	"net/http"
)

func POST[R any](g *gin.RouterGroup, path string, req *R, handler func(*gin.Context, *R) (any, error)) {
	g.POST(path, func(c *gin.Context) {
		handle(c, req, handler)
	})
}

func GET[R any](g *gin.RouterGroup, path string, req *R, handler func(*gin.Context, *R) (any, error)) {
	g.GET(path, func(c *gin.Context) {
		handle(c, req, handler)
	})
}

func DELETE[R any](g *gin.RouterGroup, path string, req *R, handler func(*gin.Context, *R) (any, error)) {
	g.DELETE(path, func(c *gin.Context) {
		handle(c, req, handler)
	})
}

func handle[R any](c *gin.Context, r R, f func(c *gin.Context, r R) (any, error)) {
	if err := c.ShouldBind(&r); err != nil {
		c.JSON(http.StatusOK, callback.BackFail("参数错误"))
		return
	}
	v, err := f(c, r)
	if err != nil {
		c.JSON(http.StatusOK, callback.BackFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, callback.SuccessData(v))
}

func SetupRouter(r *gin.Engine) {
	r.GET("/ping", api.Health)
	r.GET("/", api.UI)
	rg := r.Group("/")
	POST(rg, "register", &req.Register{}, api.User.Register)
	POST(rg, "login", &req.Login{}, api.User.Login)
	r.GET("/im", api.Im.Conn)
}
