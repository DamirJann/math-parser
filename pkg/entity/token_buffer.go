package entity

type TokenBuffer interface {
	NextToken() *Token
}

func (t *tokenBuffer) NextToken() *Token {
	if t.pos >= len(t.data)-1 {
		return nil
	}
	res := t.data[t.pos]
	t.pos += 1
	return res
}

type tokenBuffer struct {
	data    []*Token
	current *Token
	pos     int
}

func NewTokenBuffer(t []*Token) *tokenBuffer {
	return &tokenBuffer{
		data:    t,
		current: nil,
		pos:     0,
	}
}
