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

func (ri *RIndex) add(id uint, s string) {
	for i := 0; i <= len(s)-ri.size; i++ {
		tok := s[i : i+ri.size]
		if ri.data[tok] == nil {
			ri.data[tok] = NewBitSet().Add(id)
		} else {
			ri.data[tok].Add(id)
		}
	}
}

func (ri *RIndex) search(s string) []uint {
	var result *BitSet = nil

	for i := 0; i <= len(s)-ri.size; i++ {
		tok := s[i : i+ri.size]
		if ri.data[tok] == nil {
			return nil
		}
		if result == nil {
			result = NewBitSet().OR(ri.data[tok])
		} else {
			result.AND(ri.data[tok])
		}

	}

	if result == nil {
		return []uint{}
	} else {
		return result.ToArray()
	}
}
