package si

type Store struct {
	data []string
}

func NewStore() *Store {
	return &Store{
		data: make([]string, 0),
	}
}
func (store *Store) save(s string) uint {
	store.data = append(store.data, s)
	return uint(len(store.data) - 1)
}

func (store *Store) load(id uint) string {
	return store.data[id]
}
