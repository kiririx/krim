package callback

import (
	"github.com/kiririx/krim/constx"
	"github.com/kiririx/krim/module/resp"
)

func SuccessData(data any) resp.Resp {
	if data == nil {
		data = map[string]interface{}{}
	}
	result := &resp.Resp{
		Status: constx.RespSuccessStr,
		Code:   constx.RespSuccess,
		Data:   data,
	}
	return *result
}

func Error(code int, msg string) resp.Resp {
	result := resp.Resp{
		Status: constx.RespFailStr,
		ErrMsg: msg,
		Code:   code,
		Data:   map[string]interface{}{},
	}
	return result
}

func Success() resp.Resp {
	result := &resp.Resp{
		Status: constx.RespSuccessStr,
		Code:   constx.RespSuccess,
		Data:   map[string]interface{}{},
	}
	return *result
}

func BackFail(msg string) resp.Resp {
	result := &resp.Resp{
		Status: constx.RespFailStr,
		Code:   constx.RespFail,
		ErrMsg: msg,
	}
	return *result
}
