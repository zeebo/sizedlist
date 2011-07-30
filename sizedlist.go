package sizedlist

type Lenable interface {
	Len() int
}

type SizedList struct {
	Data []Lenable

	bytes int
	count int

	byte_cache int
}

func NewCountList(count int) *SizedList {
	return &SizedList{
		Data:  make([]Lenable, 0),
		count: count,
	}
}

func NewBytesList(bytes int) *SizedList {
	return &SizedList{
		Data:  make([]Lenable, 0),
		bytes: bytes,
	}
}

func NewSizedList(bytes, count int) *SizedList {
	return &SizedList{
		Data:  make([]Lenable, 0),
		bytes: bytes,
		count: count,
	}
}

//Appends items 
func (s *SizedList) Append(items ...Lenable) {
	for _, i := range items {
		s.byte_cache += i.Len()
	}
	s.Data = append(s.Data, items...)
	s.enforce()
}

//Returns the number of bytes stored in it
func (s SizedList) Len() int {
	return s.byte_cache
}

func (s *SizedList) Clear() {
	s.byte_cache = 0
	s.Data = make([]Lenable, 0)
}

func (s *SizedList) enforce() {
	count := len(s.Data)

	//enforce count
	if s.count > 0 && count > s.count {
		for _, i := range s.Data[:count-s.count] {
			s.byte_cache -= i.Len()
		}
		s.Data = s.Data[count-s.count:]
	}

	//enforce bytes
	for s.bytes > 0 && s.byte_cache > s.bytes {
		s.byte_cache -= s.Data[0].Len()
		s.Data = s.Data[1:]
	}
}
