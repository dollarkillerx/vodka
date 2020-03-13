/**
*@program: vodka
*@description: file处理相关
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-03-05 17:29
 */
package utils

import "os"

// 判断给定路径文件或在文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
