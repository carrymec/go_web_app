package tool

import (
	"fmt"
	"testing"
)

func TestPwd(t *testing.T) {
	encrypt, _ := AesEncrypt("11111")
	fmt.Println(encrypt)

	println(AesDecrypt("$2a$10$gVcykywqclr.WH6CUUNtyuMZ4qxbR7zMCVFSMOCdB2PcckQII5Mau",
		"11111"))
}
