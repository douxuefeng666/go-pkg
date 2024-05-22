/*
 * @Author: i@douxuefeng.cn
 * @Date: 2024-05-21 15:23:20
 * @LastEditTime: 2024-05-21 15:26:43
 * @LastEditors: i@douxuefeng.cn
 * @Description:
 */
package pkg

import "os"

type file struct {
}

var NewFile = newFile()

func newFile() *file {
	return &file{}
}

func (f *file) IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func (f *file) CreateFile(path string) error {
	if !f.IsExist(path) {
		return os.MkdirAll(path, os.ModePerm)
	}
	return nil
}
