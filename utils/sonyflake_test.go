/**
 * @Author: DollarKillerX
 * @Description: sonyflake_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午8:49 2020/1/3
 */
package utils

import (
	"log"
	"testing"
)

func TestSonyFlakeGetId(t *testing.T) {
	s, e := SonyFlakeGetId()
	if e != nil {
		log.Fatalln(e)
	}
	log.Println(s)
}

