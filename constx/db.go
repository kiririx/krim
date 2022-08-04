package constx

import "errors"

var (
	DBSaveFail       = 0
	DBRecordNotFound = errors.New("用户不存在")
)
