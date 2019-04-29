package common

type UintHash interface {
	Hash(string) uint32
}

type BloomFilter struct {
	bm    *BitMap
	funcs []UintHash
}

func NewDefaultBloomFilter() *BloomFilter {
	var funcs []UintHash
	for i := 0; i < 6; i++ {
		s := SimpleHash{
			Cap:  1 << (32 - 1),
			Seed: uint32(i),
		}
		funcs = append(funcs, s)
	}
	return &BloomFilter{
		bm:    NewBitMap(1 << 32),
		funcs: funcs,
	}
}

func NewBloomFilter(funcs []UintHash) *BloomFilter {
	return &BloomFilter{
		bm:    NewBitMap(1 << 32),
		funcs: funcs,
	}
}

func (b *BloomFilter) Add(k string) {
	for _, f := range b.funcs {
		pos := f.Hash(k)
		b.bm.SetBit(uint64(pos), 1)
	}
}

func (b BloomFilter) Has(k string) bool {
	hasCount := 0
	for _, f := range b.funcs {
		pos := f.Hash(k)
		if b.bm.GetBit(uint64(pos)) > 0 {
			hasCount++
		}
	}
	if hasCount == len(b.funcs) {
		return true
	}
	return false
}

type SimpleHash struct {
	Cap  uint32
	Seed uint32
}

func (s SimpleHash) Hash(k string) uint32 {
	var result uint32 = 0
	for i := 0; i < len(k); i++ {
		result = result*s.Seed + uint32(k[i])
	}
	return (s.Cap - 1) & result
}
