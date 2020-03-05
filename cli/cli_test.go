/**
*@program: vodka
*@description: https://github.com/dollarkillerx
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-03-05 18:04
 */
package main

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestProtoPath(t *testing.T) {
	join := filepath.Join("a", "generate")
	fmt.Println(join)
}
