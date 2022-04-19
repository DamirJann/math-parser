package entity

type Stack interface {
	Pop() Token
	Push(Token)
	Data() []Token
	Size() int
}

type stack struct {
	data []Token
}

func NewStack() Stack {
	return &stack{}
}

func (s stack) Size() int {
	return len(s.Data())
}

func (s stack) Data() []Token {
	return s.data
}

func (s *stack) Pop() Token {
	pop := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return pop
}

func (s *stack) Push(t Token) {
	s.data = append(s.data, t)
}
