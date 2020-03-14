/**
*@program: vodka
*@description: env相关
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-03-13 20:19
 */
package env

import "os"

const (
	VODKA_ENV   = "VODKA_ENV"
	PRODUCT_ENV = "product"
	TEST_ENV    = "test"
)

var (
	env = TEST_ENV
)

func init() {
	nenv := os.Getenv(VODKA_ENV)
	if nenv == PRODUCT_ENV {
		env = PRODUCT_ENV
	}
}

func IsProduct() bool {
	return env == PRODUCT_ENV
}

func IsTest() bool {
	return env == TEST_ENV
}

func GetEnv() string {
	return env
}
