package common

const (
	MaxBitSize = 0x01 << 32 // 最大支持的偏移量
)

type BitMap struct {
	data []byte // 数据
	size uint64 // 最大偏移量
}

func NewBitMap(max uint64) *BitMap {
	size := max
	if max > MaxBitSize {
		size = MaxBitSize
	}
	return &BitMap{
		data: make([]byte, size>>3), // 每个key 8位，所以slice长度为最大偏移量右移3
		size: size,
	}
}

func (b *BitMap) SetBit(offset uint64, val uint8) bool {
	index, pos := offset/8, offset%8
	if offset >= b.size {
		return false
	}
	if val == 0 {
		b.data[index] &^= 0x01 << pos
	} else {
		b.data[index] |= 0x01 << pos
	}
	return true
}

func (b BitMap) GetBit(offset uint64) uint8 {
	index, pos := offset/8, offset%8
	if offset >= b.size {
		return 0
	}
	return (b.data[index] >> pos) & 0x01
}
