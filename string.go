package tool

import (
	"strings"
)

// ToCamel 下划线转驼峰
func ToCamel(str string, capitalize bool) (camelStr string) {
	if len(str) == 0 {
		return
	}

	for k, v := range strings.Split(str, "_") {
		// 首字母是否大写
		if k == 0 && !capitalize {
			camelStr += v
		} else {
			camelStr += strings.ToUpper(string(v[0])) + v[1:]
		}
	}

	return
}
