package si

type SearchIndex struct {
	store  *Store
	rindex *RIndex
}

func NewSearchIndex(size int) *SearchIndex {
	return &SearchIndex{
		store:  NewStore(),
		rindex: NewRIndex(size),
	}
}
func (si *SearchIndex) add(s string) bool {
	id := si.store.save(s)
	si.rindex.add(id, s)
	return true
}

func (si *SearchIndex) search(s string) []string {
	docs := []string{}
	ids := si.rindex.search(s)
	for i := 0; i < len(ids); i++ {
		docs = append(docs, si.store.load(ids[i]))
	}
	return docs
}
