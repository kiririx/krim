package sqlx

import (
	"context"
	"database/sql"
	"errors"
	"github.com/kiririx/krim/conf"
	"gorm.io/gorm"
	"sync"
)

var (
	Ctl  = conf.Sqlx
	Conn *sql.DB
)

type Tx struct {
	once   sync.Once
	commit bool
	Sql    *gorm.DB
	Force  bool // 强制提交
}

func (t *Tx) Terminate() {
	t.once.Do(func() {
		if t.Force == true {
			t.Sql.Commit()
			return
		}

		if t.commit {
			t.Sql.Commit()
			return
		}

		t.Sql.Rollback()

		return
	})
}

func (t *Tx) Commit() {
	t.commit = true
}

func (t *Tx) Fail() {
	t.commit = false
}

func (t *Tx) Check() bool {
	return t.commit
}

func (t *Tx) Error(m string) error {
	t.commit = false
	return errors.New(m)
}

func (t *Tx) Create(data interface{}) error {
	return t.Sql.Create(data).Error
}

func (t *Tx) Save(data interface{}) error {
	return t.Sql.Save(data).Error
}

func (t *Tx) Update(model interface{}, data interface{}) error {
	return t.Sql.Model(model).Updates(data).Error
}

func Transaction() *Tx {
	return &Tx{Sql: Ctl.Begin(), commit: false, Force: false}
}

func GetTx(ctx context.Context) *gorm.DB {
	tx := ctx.Value("_tx")
	if tx != nil {
		return tx.(*Tx).Sql
	}
	return Ctl
}

func RequireGetTx(ctx context.Context) {
}
func CreateTx(ctx context.Context) context.Context {
	tx := Transaction()
	return context.WithValue(ctx, "_tx", tx)
}
