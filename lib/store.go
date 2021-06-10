package si

type Store interface {
	Save(s string) uint
	Load(id uint) string
}

type MemoryStore struct {
	data []string
}

func NewStore() Store {
	return &MemoryStore{
		data: make([]string, 0),
	}
}
func (store *MemoryStore) Save(s string) uint {
	store.data = append(store.data, s)
	return uint(len(store.data) - 1)
}

func (store *MemoryStore) Load(id uint) string {
	return store.data[id]
}
