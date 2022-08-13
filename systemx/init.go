package systemx

import (
	"github.com/kiririx/krim/service"
)

func init() {
	_, err := service.UserService.ReRegister("#contact_validator", "好友验证", "0")
	if err != nil {
		panic("system init fail")
	}
}
