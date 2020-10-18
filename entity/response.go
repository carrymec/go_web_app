package entity

type JsonResult struct {
	Code int8        `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
