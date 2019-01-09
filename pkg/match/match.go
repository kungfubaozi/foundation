package fs_regx_match

import (
	"regexp"
)

func Phone(str string) bool {
	reg := `^1([38][0-9]|14[57]|5[^4])\d{8}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(str)
}

func Email(str string) bool {
	reg := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(str)
}
