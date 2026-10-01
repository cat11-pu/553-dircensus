package store

// Store keeps records.
type Store struct {
	items []string
}

/*
multi line comment
closed below
*/
func (s *Store) Add(item string) {
	s.items = append(s.items, item)
}
