package utils

import (
	"bufio"
	"io/ioutil"
	"os"
)

// FileGetContents 获取文件内容
func FileGetContents(filename string) ([]byte, error) {
	fp, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)

	if err != nil {
		return nil, err
	}
	defer fp.Close()
	reader := bufio.NewReader(fp)
	return ioutil.ReadAll(reader)
}

// FilePutContents 存储内容到文件
func FilePutContents(content string, filePath string) error {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer f.Close()
	if err != nil {
		return err
	}
	return nil
}
