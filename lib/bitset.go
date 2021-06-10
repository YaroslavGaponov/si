package si

import (
	"math"
	"math/bits"
)

var ARCH_BITS = bits.Len(bits.UintSize)
var ARCH_MASK = uint(math.Pow(2, 7)) - 1

type BitSet struct {
	data map[uint]uint
}

func NewBitSet() *BitSet {
	return &BitSet{
		data: make(map[uint]uint),
	}
}

func (bs *BitSet) Add(value uint) *BitSet {
	base := value >> ARCH_BITS
	offset := value & ARCH_MASK
	bs.data[base] |= 1 << offset
	return bs
}

func (bs *BitSet) Remove(value uint) *BitSet {
	base := value >> ARCH_BITS
	offset := value & ARCH_MASK
	bs.data[base] &= ^(1 << offset)
	return bs
}

func (bs *BitSet) Has(value uint) bool {
	base := value >> ARCH_BITS
	offset := value & ARCH_MASK
	return bs.data[base]&(1<<offset) == 1
}

func (bs *BitSet) ToArray() []uint {
	result := []uint{}
	for base, offset := range bs.data {
		for i := 0; i < ARCH_BITS; i++ {
			if offset&uint(1<<i) != 0 {
				result = append(result, uint(base<<ARCH_BITS|uint(i)))
			}
		}
	}
	return result
}

func (bs *BitSet) AND(bs2 *BitSet) *BitSet {
	for base, offset := range bs2.data {
		bs.data[base] &= offset
		if bs.data[base] == 0 {
			delete(bs.data, base)
		}
	}
	return bs
}

func (bs *BitSet) OR(bs2 *BitSet) *BitSet {
	for base, offset := range bs2.data {
		bs.data[base] |= offset
		if bs.data[base] == 0 {
			delete(bs.data, base)
		}
	}
	return bs
}
