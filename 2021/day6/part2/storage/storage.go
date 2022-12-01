package storage

type Storage struct {
	list [9]int
}

// Make a iteration step
func (s *Storage) Next() {
	zeroVal := s.list[0]

	for i := 1; i < len(s.list); i++ {
		s.list[i-1] = s.list[i]
	}
	s.list[8] = 0

	if zeroVal > 0 {
		s.list[8] += zeroVal
		s.list[6] += zeroVal
	}
}

// Sum all values from storage list
func (s *Storage) Summary() int {
	sum := 0

	for _, val := range s.list {
		sum += val
	}

	return sum
}

// Create and fill storage
func Create(list []int) Storage {
	store := Storage{}

	for _, val := range list {
		store.list[val]++
	}

	return store
}
