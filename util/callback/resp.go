package callback

import (
	"github.com/kiririx/krim/constant"
	"github.com/kiririx/krim/module/resp"
)

func SuccessData(data any) resp.Resp {
	if data == nil {
		data = map[string]interface{}{}
	}
	result := &resp.Resp{
		Status: constant.RespSuccessStr,
		Code:   constant.RespSuccess,
		Data:   data,
	}
	return *result
}

func Error(code int, msg string) resp.Resp {
	result := resp.Resp{
		Status: constant.RespFailStr,
		ErrMsg: msg,
		Code:   code,
		Data:   map[string]interface{}{},
	}
	return result
}

func Success() resp.Resp {
	result := &resp.Resp{
		Status: constant.RespSuccessStr,
		Code:   constant.RespSuccess,
		Data:   map[string]interface{}{},
	}
	return *result
}

func BackFail(msg string) resp.Resp {
	result := &resp.Resp{
		Status: constant.RespFailStr,
		Code:   constant.RespFail,
		ErrMsg: msg,
	}
	return *result
}
