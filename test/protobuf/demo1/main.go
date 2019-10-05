/**
 * @Author: DollarKiller
 * @Description: protobuf test
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 18:36 2019-10-03
 */
package main

import (
	"github.com/dollarkillerx/vodka/test/protobuf/demo1/phone"
	"github.com/gogo/protobuf/proto"
	"io/ioutil"
	"log"
)

func main() {
	var person phone.Person
	person.Name = "DollarKiller"
	person.Id = 232324
	var phonez phone.Phone
	phonez.Number = "123213123"
	person.Phones = append(person.Phones, &phonez)

	bytes, e := proto.Marshal(&person)
	if e != nil {
		panic(e)
	}

	e = ioutil.WriteFile("test/protobuf/testfile", bytes, 00666)
	if e != nil {
		panic(e)
	}

	file, e := ioutil.ReadFile("test/protobuf/testfile")
	if e != nil {
		panic(e)
	}
	var dataz phone.Person
	e = proto.Unmarshal(file, &dataz)
	if e != nil {
		panic(e)
	}

	log.Println(dataz)

}
