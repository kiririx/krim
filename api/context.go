package api

import "github.com/gin-gonic/gin"

type APICtx struct {
	GinCtx   *gin.Context
	UserId   uint64
	UserName string
}
