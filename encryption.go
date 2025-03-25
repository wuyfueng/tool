package tool

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"sort"
)

// GenerateVerifyCode 生成校验码
func GenerateVerifyCode(serverIdList []int64) string {
	sort.Slice(serverIdList, func(i, j int) bool {
		return serverIdList[i] < serverIdList[j]
	})

	b, _ := json.Marshal(serverIdList)
	return fmt.Sprintf("%X", md5.Sum(b))
}
