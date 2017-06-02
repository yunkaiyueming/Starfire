package encrypt

import (
	"sort"
)

func GetSign(params map[string]string, secret string) string {
	strs := make([]string, 0)
	for key, _ := range params {
		strs = append(strs, key)
	}
	sort.Strings(strs)

	sign := ""
	for _, key := range strs {
		sign += key + "=" + params[key]
	}
	sign += secret
	return Md5(sign)
}

func CheckSign(params map[string]string, secret, resSign string) bool {
	expectedSign := GetSign(params, secret)
	if resSign == expectedSign {
		return true
	}
	return false
}
