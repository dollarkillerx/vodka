/**
 * @Author: DollarKiller
 * @Description: utils str 工具测试
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:47 2019-10-04
 */
package utils

import (
	"github.com/dollarkillerx/vodka/utils"
	"os"
	"testing"
)

func TestPath(t *testing.T) {
	slash := utils.PathSlash("hello/")
	t.Log(slash)

	slash = utils.PathSlash("hello")
	t.Log(slash)
}

func TestMkdir(t *testing.T) {
	all := os.MkdirAll("./output/controllerx", 00666)
	if all != nil {
		panic(all)
	}
}
