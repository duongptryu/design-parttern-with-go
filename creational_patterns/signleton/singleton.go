package signleton

type Signleton interface {
	AddOne() int
	GetCount() int
}

type singleton struct {
	count int
}

var instance *singleton

func GetInstance() Signleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count += 1
	return s.count
}

func (s *singleton) GetCount() int {
	return s.count
}
