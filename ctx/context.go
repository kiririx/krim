package ctx

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/kiririx/krim/sqlx"
	"github.com/kiririx/krutils/logx"
	"gorm.io/gorm"
)

type Ctx struct {
	GinCtx   *gin.Context
	UserId   uint64
	UserName string
	Tx       *sqlx.Tx
}

func (c *Ctx) SqlCtl() *gorm.DB {
	if c.Tx != nil {
		return c.Tx.Sql
	}
	return sqlx.Ctl
}

func (c *Ctx) Fail(message string) {
	if c.Tx != nil {
		c.Tx.Fail()
	}
	logx.ERR(errors.New(message))
}

func (c *Ctx) Finish() {
	if c.Tx != nil {
		c.Tx.Terminate()
	}
}

func (c *Ctx) CommitTx() {
	c.Tx.Commit()
}

func (c *Ctx) CreateTx() {
	c.Tx = sqlx.Transaction()
}
