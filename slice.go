package tool

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
)

const (
	Separator = ";" // 默认分隔符
)

// StringToMapInt32 字符串转map
func StringToMapInt32(str string) (m map[int32]bool) {
	m = make(map[int32]bool, 1)
	for _, v := range strings.Split(str, Separator) {
		id, err := strconv.Atoi(v)
		if err == nil {
			m[int32(id)] = true
		}
	}
	return
}

// StringToInt32 string类型的字符串转换为int32类型的切片
func StringToInt32(str string) (ids []int32) {
	for _, v := range strings.Split(str, Separator) {
		id, err := strconv.Atoi(v)
		if err == nil {
			ids = append(ids, int32(id))
		}
	}
	return
}

// StringToInt64 string类型的字符串转换为int64类型的切片
func StringToInt64(str string) (ids []int64) {
	for _, v := range strings.Split(str, Separator) {
		id, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
			ids = append(ids, id)
		}
	}
	return
}

// SliceIntToStr int切片转字符串
func SliceIntToStr(ids []int) (str string) {
	for _, v := range ids {
		str = StrUnaryAppend(str, v)
	}

	return
}

// SliceInt32ToStr int32切片转字符串
func SliceInt32ToStr(ids []int32) (str string) {
	for _, v := range ids {
		str = StrUnaryAppend(str, v)
	}

	return
}

// SliceInt64ToStr int64切片转字符串
func SliceInt64ToStr(ids []int64) (str string) {
	for _, v := range ids {
		str = StrUnaryAppend(str, v)
	}

	return
}

// IsStrUnaryExist 字符串一维数组 是否存在某个元素
func IsStrUnaryExist(str string, elem interface{}) bool {
	for _, v := range strings.Split(str, Separator) {
		if v == fmt.Sprintf("%v", elem) {
			return true
		}
	}

	return false
}

// StrUnaryAppend 字符串一维数组 追加元素
func StrUnaryAppend(str string, elem interface{}) string {
	if str == "" {
		return fmt.Sprintf("%v", elem)
	} else {
		return fmt.Sprintf("%s%s%v", str, Separator, elem)
	}
}

// StrUnaryRemove 字符串一维数组 移除元素
func StrUnaryRemove(str, elem string) string {
	list := strings.Split(str, Separator)
	for k, v := range list {
		if v == elem {
			list = append(list[:k], list[k+1:]...)
			break
		}
	}

	return strings.Join(list, Separator)
}

// ShuffleSlice 打乱切片
func ShuffleSlice(x interface{}) {
	rv := reflect.ValueOf(x)
	swap := reflect.Swapper(x)
	length := rv.Len()
	for i := length - 1; i > 0; i-- {
		// 在i和0之间选择一个随机数
		swap(i, rand.Intn(i+1))
	}
}

// InArrayInt32 存在于[]int32
func InArrayInt32(arr []int32, elem int32) bool {
	for _, v := range arr {
		if v == elem {
			return true
		}
	}
	return false
}

// InArrayInt64 存在于[]int64
func InArrayInt64(arr []int64, elem int64) bool {
	for _, v := range arr {
		if v == elem {
			return true
		}
	}
	return false
}

// StringToSS 字符串自定义转字符串切片
func StringToSS(str, separator string, f func(str string) string) (ss []string) {
	for _, v := range strings.Split(str, separator) {
		ss = append(ss, f(v))
	}
	return
}
