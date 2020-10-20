package param

type SmsLogin struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}
