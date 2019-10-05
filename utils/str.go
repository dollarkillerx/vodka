/**
 * @Author: DollarKiller
 * @Description: str处理
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:45 2019-10-04
 */
package utils

// 如果path 存在 "/"就去掉
func PathSlash(path string) string {
	if string(path[len(path)-1]) == "/" {
		path = path[:len(path)-1]
	}

	return path
}
