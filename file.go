package tool

import "os"

// AppendFile 类似于 os.WriteFile，但支持追加模式
func AppendFile(filename string, data []byte, perm os.FileMode) error {
	// 打开文件以追加模式
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, perm)
	if err != nil {
		return err
	}
	defer file.Close()

	// 写入数据
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
