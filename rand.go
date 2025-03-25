package tool

import (
	"errors"
	"math/rand"
)

// GetRandNumByMinmax 获取随机数 根据最大,最小值
func GetRandNumByMinmax(min, max int64) (num int64, err error) {
	if min < 0 || min > max {
		err = errors.New("param invalid")
		return
	}

	num = rand.Int63n(max-min+1) + min

	return
}

// GetSpecialRandNumByMinmax 获取随机数 根据最大,最小值
func GetSpecialRandNumByMinmax(r *rand.Rand, min, max int64) (num int64, err error) {
	if min < 0 || min > max {
		err = errors.New("param invalid")
		return
	}

	num = r.Int63n(max-min+1) + min

	return
}

// RandNums 获取[0,n)中不重复的x个数
// params[0] 遍历次数上限
func RandNums(n int, x int, params ...int) (nums map[int]bool) {
	if n <= 0 || x <= 0 {
		return
	}

	// 遍历次数上限为1000
	eachNum := 1000
	if len(params) == 1 {
		eachNum = params[0]
	}

	if x > n {
		x = n
	}

	nums = make(map[int]bool, x)
	for i := 0; i < eachNum; i++ {
		if len(nums) >= x {
			break
		}

		randNum := rand.Intn(n)
		if !nums[randNum] {
			nums[randNum] = true
		}
	}

	return
}

// SpecialRandNums 获取[0,n)中不重复的x个数
// params[0] 遍历次数上限
func SpecialRandNums(r *rand.Rand, n int, x int, params ...int) (nums map[int]bool) {
	if n <= 0 || x <= 0 {
		return
	}

	// 遍历次数上限为1000
	eachNum := 1000
	if len(params) == 1 {
		eachNum = params[0]
	}

	if x > n {
		x = n
	}

	nums = make(map[int]bool, x)
	for i := 0; i < eachNum; i++ {
		if len(nums) >= x {
			break
		}

		randNum := r.Intn(n)
		if !nums[randNum] {
			nums[randNum] = true
		}
	}

	return
}

// RandRangeInt 获取随机数int 根据最大,最小值
func RandRangeInt(min, max int) int {
	if min > max {
		return 0
	}
	return rand.Intn(max-min+1) + min
}
