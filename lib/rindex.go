package si

type RIndex struct {
	size int
	data map[string]*BitSet
}

func NewRIndex(size int) *RIndex {
	return &RIndex{
		size: size,
		data: make(map[string]*BitSet),
	}
}

func (ri *RIndex) Add(id uint, s string) {
	for i := 0; i <= len(s)-ri.size; i++ {
		tok := s[i : i+ri.size]
		if ri.data[tok] == nil {
			ri.data[tok] = NewBitSet().Add(id)
		} else {
			ri.data[tok].Add(id)
		}
	}
}

func (ri *RIndex) searchOne(s string) *BitSet {
	result := NewBitSet()
	for i := 0; i <= len(s)-ri.size; i++ {
		tok := s[i : i+ri.size]
		if ri.data[tok] == nil {
			return nil
		}
		if i == 0 {
			result.OR(ri.data[tok])
		} else {
			result.AND(ri.data[tok])
		}

	}
	return result
}

func (ri *RIndex) Search(s []string) []uint {
	var result *BitSet

	for i, word := range s {
		if i == 0 {
			result = ri.searchOne(word)
		} else {
			result.AND(ri.searchOne(word))
		}
	}

	return result.ToArray()
}
