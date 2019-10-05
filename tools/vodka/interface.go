/**
 * @Author: DollarKiller
 * @Description: cli interface
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:18 2019-10-04
 */
package main

type Generator interface {
	Run(opt *Option) error
}
