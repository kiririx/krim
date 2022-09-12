package systemx

import (
	"github.com/kiririx/krim/logic"
	"github.com/kiririx/krutils/logx"
	"runtime/debug"
)

func init() {
	_, err := logic.UserLogic.ReRegister("#contact_validator", "好友验证", "0")
	if err != nil {
		logx.ERR(err)
		debug.PrintStack()
		panic("system init fail")
	}
}
