package mail

import (
	"testing"
)

func TestSendMessageCode(t *testing.T) {
	//regionId := "ap-northeast-1"
	//accessKeyId := ""
	//accessKeySecret := ""
	//signName := ""
	//tempCode := ""
	err := SendMessageCode("123455", "16278930002")
	if err != nil {
		panic(err)
	}
}
