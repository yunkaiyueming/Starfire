package encrypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSign(t *testing.T) {
	type Input struct {
		params map[string]string
		secret string
	}

	type Out struct {
		ret string
	}

	inputs := []Input{
		{map[string]string{"app_name": "bi", "user": "zhangsan@rayjoy.com", "url": "http://www.google.com"}, "723773045578c74442ace5af8f3b0e8359a4c094"},
		{map[string]string{"app_name": "hammer", "user": "lisi@rayjoy.com", "url": "http://www.baidu.com"}, "d87c93907550081d61cb673d6d7eba77312d8bf3"},
		{map[string]string{"app_name": "rsdk-set", "user": "longlong@rayjoy.com", "url": "http://www.hao123.com"}, "bc4d7d2e64c6fb343dd31a8a4ece36b9a9141a28"},
	}

	shoulds := []Out{
		{"f3c99ce3f7a1be009c985ce12702fdb1"},
		{"c06a9b47a94755fc4cf2eaae95c29f64"},
		{"cd1f092c737658f18857b23441357f75"},
	}

	for k, input := range inputs {
		ret := GetSign(input.params, input.secret)
		assert.Equal(t, shoulds[k].ret, ret)
	}
}

func TestCheckSign(t *testing.T) {
	type Input struct {
		params map[string]string
		secret string
		sign   string
	}

	type Out struct {
		ret bool
	}

	inputs := []Input{
		{map[string]string{"app_name": "bi", "user": "zhangsan@rayjoy.com", "url": "http://www.google.com"}, "723773045578c74442ace5af8f3b0e8359a4c094", "f3c99ce3f7a1be009c985ce12702fdb1"},
		{map[string]string{"app_name": "hammer", "user": "lisi@rayjoy.com", "url": "http://www.baidu.com"}, "d87c93907550081d61cb673d6d7eba77312d8bf3", "c06a9b47a94755fc4cf2eaae95c29f64"},
		{map[string]string{"app_name": "rsdk-set", "user": "longlong@rayjoy.com", "url": "http://www.hao123.com"}, "d87c93907550081d61cb673d6d7eba77312d8bf3", "ttt"},
	}

	shoulds := []Out{
		{true},
		{true},
		{false},
	}

	for k, input := range inputs {
		ret := CheckSign(input.params, input.secret, input.sign)
		assert.Equal(t, shoulds[k].ret, ret)
	}
}
