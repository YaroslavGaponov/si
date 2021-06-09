package si

type Store struct {
	data []string
}

func NewStore() *Store {
	return &Store{
		data: make([]string, 0),
	}
}
func (store *Store) Save(s string) uint {
	store.data = append(store.data, s)
	return uint(len(store.data) - 1)
}

func (store *Store) Load(id uint) string {
	return store.data[id]
}
