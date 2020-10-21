package param

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Id       string `json:"id"`   //验证码ID
	Code     string `json:"code"` //验证码
}
