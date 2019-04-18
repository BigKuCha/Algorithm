package common

import (
	"fmt"
	"time"
)

var (
	nodeBits uint8 = 10 // 节点id的位数， 总共能表示1024个节点
	stepBits uint8 = 12 // 序列号的位数，统一毫秒能产生1<<12=4096个不重复id

	nodeMax int64 = 1<<nodeBits - 1 // 最大支持的节点编号
	stepMax int64 = 1<<stepBits - 1 // 最大支持的序列号

	timeShift uint8 = nodeBits + stepBits // 时间戳向左偏移量
	nodeShift uint8 = stepBits            // 节点id向左偏移量

	lastTimestamp int64     // 上次的时间戳
	sn            int64 = 0 // 序列号
)

func init() {
	lastTimestamp = time.Now().UnixNano() / 1e6
}

func Snowflake(id int64) int64 {
	// 节点编号范围0-1023
	if id < 0 || id >= nodeMax {
		return 0
	}
	timestamp := time.Now().UnixNano() / 1e6
	if timestamp == lastTimestamp {
		sn++
		fmt.Println("时间相同了")
		if sn > stepMax {
			fmt.Println("超出最大序列号了+++")
			for timestamp <= lastTimestamp {
				timestamp = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		sn = 0
	}
	// 生成ID
	return timestamp<<timeShift | id<<nodeShift | sn
}
