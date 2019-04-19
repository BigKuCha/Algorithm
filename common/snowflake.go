package common

import (
	"sync"
	"time"
)

var (
	nodeBits uint8 = 10 // 节点id的位数， 总共能表示1024个节点
	stepBits uint8 = 12 // 序列号的位数，统一毫秒能产生1<<12=4096个不重复id

	nodeMax int64 = 1<<nodeBits - 1 // 最大支持的节点编号
	stepMax int64 = 1<<stepBits - 1 // 最大支持的序列号

	timeShift uint8 = nodeBits + stepBits // 时间戳向左偏移量
	nodeShift uint8 = stepBits            // 节点id向左偏移量

)

type SnowFlake struct {
	machineID int64
	lastTime  int64
	sn        int64
	sync.Mutex
}

func NewSnowFlake(machineID int64) *SnowFlake {
	if machineID < 0 || machineID > nodeMax {
		return nil
	}
	return &SnowFlake{
		machineID: machineID,
		lastTime:  time.Now().UnixNano() / 1e6,
	}
}

func (s *SnowFlake) GetID() int64 {
	s.Lock()
	defer s.Unlock()
	now := time.Now().UnixNano() / 1e6
	if now == s.lastTime {
		s.sn++
		if s.sn > stepMax {
			for now <= s.lastTime {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		s.sn = 0
	}
	s.lastTime = now
	return now<<timeShift | s.machineID<<nodeShift | s.sn
}
