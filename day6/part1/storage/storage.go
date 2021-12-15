package storage

type lanternfish struct {
	created string
}

type storage struct {
	list [][]lanternfish
}

func (s *storage) insert(value int, el lanternfish) {
	int8Val := int8(value)
	(*s).list[int8Val] = append((*s).list[int8(value)], el)
}

func (s *storage) Next(creation string) {
	zeroArr := s.list[0]

	for i := 1; i < len(s.list); i++ {
		s.list[i-1] = s.list[i]
	}
	s.list[8] = make([]lanternfish, 0)

	if len(zeroArr) > 0 {
		for _, lanterfishEl := range zeroArr {
			s.list[8] = append(s.list[8], lanternfish{
				created: creation,
			})
			s.list[6] = append(s.list[6], lanterfishEl)
		}
	}
}

func (s *storage) Summary() int {
	sum := 0

	for i := 0; i < len(s.list); i++ {
		sum += len(s.list[i])
	}

	return sum
}

// Create new instance of storage with init values
func Create(values []int) storage {
	store := storage{
		list: make([][]lanternfish, 9),
	}

	for i := int8(0); i <= 8; i++ {
		store.list[i] = make([]lanternfish, 0)
	}

	for _, el := range values {
		store.list[el] = append(store.list[el], lanternfish{
			created: "Init",
		})
	}

	return store
}
