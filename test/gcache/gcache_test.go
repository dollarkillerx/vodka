/**
 * @Author: DollarKiller
 * @Description: gcache test
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 10:38 2019-10-07
 */
package gcache

import (
	"github.com/bluele/gcache"
	"log"
	"testing"
)

func TestGet(t *testing.T) {
	cache := gcache.New(20).LRU().Build()

	cache.Set("ok","asd")
	cache.Set("ok1","asd11")
	cache.Set("ok2","asd22")

	//get, e := cache.Get("ok")
	//if e != nil {
	//	panic(e)
	//}else {
	//	log.Println(get)
	//}

	all := cache.GetALL(true)
	for _,i := range all {
		log.Println(i)
	}
}
