/**
*@program: vodka
*@description: https://github.com/dollarkillerx
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-03-06 19:39
 */
package template_test

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"testing"
	"time"
)

var tmp1 = `
{{if and (eq .Ag false) (eq .Pc false)}}
	OOK
{{else}}
	false
{{end}}
`

var tmp2 = `

`

var tmp3 = `

`

func TestTmp(t *testing.T) {
	t2 := template.New("aa")
	parse, err := t2.Parse(tmp1)
	if err != nil {
		log.Fatalln(err)
	}

	//i := make([]byte, 2048)
	//buffer := bytes.NewBuffer(i)
	//buffer.WriteString("aaa")
	//readString, err := buffer
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Println(readString)
	////create, err := os.Create("a")
	////if err != nil {
	////	log.Fatalln(err)
	////}
	a := ""
	bufferString := bytes.NewBufferString(a)
	err = parse.Execute(bufferString, map[string]interface{}{
		"Ag": false,
		"Pc": false,
	})


	if err != nil {
		log.Fatalln(err)
	}
	time.Sleep(time.Second)
	fmt.Println(bufferString.String())
}

func TestWW(t *testing.T) {
	a := ""
	bufferString := bytes.NewBufferString(a)
	_, err := bufferString.WriteString("aaa")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(bufferString.String())
}