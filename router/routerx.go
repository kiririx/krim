package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kiririx/krim/api"
	"github.com/kiririx/krim/ctx"
	"github.com/kiririx/krim/module/req"
	"github.com/kiririx/krim/util/callback"
	"github.com/kiririx/krutils/sugar"
	"net/http"
)

func buildAPICtx(c *gin.Context) *ctx.Ctx {
	return &ctx.Ctx{
		GinCtx: c,
		UserId: func() uint64 {
			userId, exists := c.Get("userId")
			return sugar.ThenFunc(exists, func() uint64 {
				return userId.(uint64)
			}, func() uint64 {
				return 0
			})
		}(),
		UserName: func() string {
			username, exists := c.Get("userName")
			return sugar.ThenFunc(exists, func() string {
				return username.(string)
			}, func() string {
				return ""
			})
		}(),
	}
}

func POST[R any](g *gin.RouterGroup, path string, req *R, handler func(*ctx.Ctx, *R) (any, error), middlewares ...gin.HandlerFunc) {
	mws := make([]gin.HandlerFunc, 0)
	mws = append(mws, middlewares...)
	mws = append(mws, func(c *gin.Context) {
		handle(buildAPICtx(c), req, handler)
	})
	g.POST(path, mws...)
}

func GET[R any](g *gin.RouterGroup, path string, req *R, handler func(*ctx.Ctx, *R) (any, error)) {
	g.GET(path, func(c *gin.Context) {
		handle(buildAPICtx(c), req, handler)
	})
}

func DELETE[R any](g *gin.RouterGroup, path string, req *R, handler func(*ctx.Ctx, *R) (any, error)) {
	g.DELETE(path, func(c *gin.Context) {
		handle(buildAPICtx(c), req, handler)
	})
}

func handle[R any](ctx *ctx.Ctx, r R, f ...func(c *ctx.Ctx, r R) (any, error)) {
	c := ctx.GinCtx
	for i, fc := range f {
		if c.Request.Method == "GET" || c.Request.Method == "DELETE" {
			if err := c.ShouldBindUri(&r); err != nil {
				c.JSON(http.StatusOK, callback.Error(0, "参数错误"))
				return
			}
		}
		if c.Request.Method == "POST" {
			if err := c.ShouldBindJSON(&r); err != nil {
				c.JSON(http.StatusOK, callback.Error(0, "参数错误"))
				return
			}
		}
		ctx.CreateTx()
		v, err := fc(ctx, r)
		if err != nil {
			ctx.Fail(err.Error())
			c.JSON(http.StatusOK, callback.Error(0, err.Error()))
			return
		}
		ctx.CommitTx()
		ctx.Finish()
		if i == len(f)-1 {
			c.JSON(http.StatusOK, callback.SuccessData(v))
		}
	}
}

func SetupRouter(r *gin.Engine) {
	r.GET("/ping", api.Health)
	r.GET("/conndtl", api.ConnDetail)
	r.GET("/", api.UI)
	rg := r.Group("/")
	POST(rg, "register", &req.Register{}, api.User.Register)
	POST(rg, "login", &req.Login{}, api.User.Login)
	r.GET("/im", api.Im.Conn)

	contact := r.Group("/contact")
	// GET(contact, "/", &req.AddContact{}, api.ContactAPI.AddContact)
	GET(contact, "/:username", &req.GetContact{}, api.ContactAPI.GetContact)
	POST(contact, "/addUser", &req.AddContactEvent{}, api.ContactAPI.AddContactEvent, api.CheckLogin)

	message := r.Group("/message")
	POST(message, "/sendUser", &req.SendUserMessageReq{}, api.SendUserMessage)
	POST(message, "/sendGroup", &req.SendGroupMessageReq{}, api.SendGroupMessage)
}
