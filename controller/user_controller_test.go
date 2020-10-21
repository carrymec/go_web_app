package controller

import (
	"go_web_app/param"
	"go_web_app/service"
	"testing"
)

func TestUserController_sendMsg(t *testing.T) {
	type args struct {
		phone string
		code  string
	}
	tests := []struct {
		name string
		args args
		ok   bool
	}{
		{
			args: args{
				phone: "",
				code:  "",
			},
			ok: false,
		},
		{
			args: args{
				phone: "17723458765",
				code:  "2222",
			},
			ok: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var userService service.UserService
			userService.FindByPhone(param.SmsLogin{
				Phone: tt.args.phone,
				Code:  tt.args.code,
			})
		})
	}

}
