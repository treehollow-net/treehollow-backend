package utils

import (
	"fmt"
	"testing"
)

var (
	cases = []struct {
		mail  string
		valid bool
	}{
		{mail: "admin@mails.tsinghua.edu.cn", valid: true},
		{mail: "thu-hole@mails.tsinghua.edu.cn", valid: true},
		{mail: "thu_hole@mails.tsinghua.edu.cn", valid: true},
		{mail: "yezhisheng@pku.edu.cn,admin@mails.tsinghua.edu.cn", valid: false},
	}
)

func TestCheckMail(t *testing.T) {
	for _, c := range cases {
		if CheckEmail(c.mail) != c.valid {
			t.Errorf("%s is expected to be %v", c.mail, c.valid)
		}
	}
}

func TestHashEmail(t *testing.T) {
	Salt = "asdfghjklpoiuyt"
	enc := HashEmail("16637378928")
	fmt.Println(enc)
}
