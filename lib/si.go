package si

type SearchIndex struct {
	store  Store
	rindex *RIndex
}

func New(size int) *SearchIndex {
	return &SearchIndex{
		store:  NewStore(),
		rindex: NewRIndex(size),
	}
}
func (si *SearchIndex) Add(s string) bool {
	id := si.store.Save(s)
	si.rindex.Add(id, s)
	return true
}

func (si *SearchIndex) Search(s string) []string {
	docs := []string{}
	ids := si.rindex.Search(s)
	for i := 0; i < len(ids); i++ {
		docs = append(docs, si.store.Load(ids[i]))
	}
	return docs
}
