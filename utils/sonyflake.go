/**
 * @Author: DollarKillerX
 * @Description: sonyflake 感谢sony公司提供伟大的分布锁id
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午4:51 2020/1/3
 */
package utils

import (
	"fmt"
	"github.com/sony/sonyflake"
	"strconv"
)

var (
	LSonyFlake *sonyflake.Sonyflake
	machineId  uint16 // 真正的分布式环境下必须zookeeper或etcd中获取
)

func getMachineID() (uint16, error) {
	return machineId, nil
}

func init() {
	Init(0)
}

func Init(mid uint16) (err error) {
	machineId = mid
	st := sonyflake.Settings{}
	st.MachineID = getMachineID
	LSonyFlake = sonyflake.NewSonyflake(st)
	return
}
func GetID() (id uint64, err error) {
	if LSonyFlake == nil {
		err = fmt.Errorf("No Init\n")
		return
	}
	return LSonyFlake.NextID()
}

func SonyFlakeGetId() (string, error) {
	id, err := GetID()
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(id)), nil
}
