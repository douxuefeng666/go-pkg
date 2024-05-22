/*
 * @Author: i@douxuefeng.cn
 * @Date: 2024-05-21 17:33:23
 * @LastEditTime: 2024-05-21 17:33:30
 * @LastEditors: i@douxuefeng.cn
 * @Description:
 */
package pkg

type JsonData struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func Json(code int, data any, msg string) *JsonData {
	return &JsonData{
		Code: code,
		Data: data,
		Msg:  msg,
	}
}

func JsonNoData(code int, msg string) *JsonData {
	return &JsonData{
		Code: code,
		Msg:  msg,
	}
}

func JsonErr(code int, err error) *JsonData {
	return &JsonData{
		Code: code,
		Msg:  err.Error(),
	}
}

func Response(code int, err error, msg string, data interface{}) *JsonData {
	if err != nil {
		return JsonErr(code, err)
	}
	return Json(200, data, msg)
}
