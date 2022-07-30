package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kiririx/krim/util/callback"
	"net/http"
)

func Health(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

// ConnDetail 连接详情
func ConnDetail(c *gin.Context) {
	c.JSON(http.StatusOK, callback.SuccessData(map[string]any{
		"conn": ClientMap,
	}))
}
