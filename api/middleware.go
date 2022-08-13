package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kiririx/krim/service"
	"github.com/kiririx/krim/util/callback"
	"net/http"
)

var (
	User = &UserApi{}
)

// GetUserId 获取用户id
func GetUserId(c *gin.Context) uint {
	if v, ok := c.Get("userId"); ok {
		return v.(uint)
	}
	return 0
}

// CheckLogin 检查是否登陆
func CheckLogin(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusOK, callback.BackFail("权限不足"))
		c.Abort()
		return
	}
	userMeta, err := service.ValidToken(token)
	if err != nil {
		c.JSON(http.StatusOK, callback.BackFail("权限不足"))
		c.Abort()
		return
	}
	c.Set("userId", userMeta.Id)
	c.Set("userName", userMeta.Username)
	c.Set("nickName", func() string {
		if userMeta != nil && userMeta.Id > 0 {
			user, err := service.UserService.QueryById(userMeta.Id)
			if err != nil {
				return ""
			}
			return user.Nickname
		}
		return ""
	}())
	c.Next()
}
