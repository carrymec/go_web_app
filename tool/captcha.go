package tool

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
	"image/color"
	"time"
)

type Captcha struct {
	Id         string `json:"id"`
	Base64Blog string `json:"base_64_blog"`
	Verity     string `json:"code"`
}

// 生成验证码
func Generate() *Captcha {
	parameters := base64Captcha.NewDriverMath(
		30,
		60,
		1,
		1,
		&color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 254,
		}, []string{})
	// 生成Id,问题,答案
	id, question, answer := parameters.GenerateIdQuestionAnswer()
	captcha, err := parameters.DrawCaptcha(question)
	if err != nil {
		return nil
	}
	// todo 将问题和答案保存到redis
	fmt.Printf("the id is %v,and the question is %v,and the answer is %v\n", id, question, answer)
	redis := GlobalRedis.client
	redis.Set(id, answer, time.Second*300)
	b64string := captcha.EncodeB64string()
	captchaResult := Captcha{
		Id:         id,
		Base64Blog: b64string,
	}
	return &captchaResult
}

func Verity(id string, code string) bool {
	redis := GlobalRedis.client
	redisAnswer, err := redis.Get(id).Result()
	if err != nil {
		return false
	}
	if redisAnswer != code {
		return false
	}
	return true
}
