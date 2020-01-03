/**
 * @Author: DollarKillerX
 * @Description: sync_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午11:18 2020/1/3
 */
package test

import (
	"log"
	"sync"
	"testing"
)

func TestSyncMap(t *testing.T) {
	i := sync.Map{}
	i.Store("ac", "0230202")
	actual, loaded := i.LoadOrStore("arc", "xxxxxxx")
	log.Println(loaded)
	log.Println(actual)

	value, ok := i.Load("ac")
	if ok {
		log.Println(value)
	}

	value, ok = i.Load("arc")
	if ok {
		log.Println(value)
	}
}
