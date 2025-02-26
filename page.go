package tools

// PageRange 分页处理
// f func(page, startIndex, endIndex int)
//      page, startIndex, endIndex: 0-index
func PageRange(count int, pageSize int, f func(page, startIndex, endIndex int)) {
	// 整页
	pageMax := count / pageSize // 最大页数
	for page := 0; page < pageMax; page++ {
		startIndex := page * pageSize
		f(page, startIndex, startIndex+pageSize-1)
	}

	// 残页(最后一页)
	if n := count % pageSize; n > 0 {
		startIndex := pageMax * pageSize
		f(pageMax, startIndex, startIndex+n-1)
	}

	return
}

// PageRangeWithScrap 分页处理(判断残页)
// f func(page, startIndex, endIndex int, isScrap bool)
//      page, startIndex, endIndex: 0-index
func PageRangeWithScrap(count int, pageSize int, f func(page, startIndex, endIndex int, isScrap bool)) {
	// 整页
	pageMax := count / pageSize // 最大页数
	for page := 0; page < pageMax; page++ {
		startIndex := page * pageSize
		f(page, startIndex, startIndex+pageSize-1, false)
	}

	// 残页(最后一页)
	if n := count % pageSize; n > 0 {
		startIndex := pageMax * pageSize
		f(pageMax, startIndex, startIndex+n-1, true)
	}

	return
}
