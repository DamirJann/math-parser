package entity

type TokenBuffer interface {
	NextToken()
	Current() *Token
	Lookahead() *Token
}

func (t *tokenBuffer) NextToken() {
	t.pos += 1
	if t.pos < len(t.data) {
		t.current = t.data[t.pos]
	} else {
		t.current = nil
	}
	if t.pos < len(t.data)-1 {
		t.lookahead = t.data[t.pos+1]
	} else {
		t.lookahead = nil
	}
}

type tokenBuffer struct {
	data      []*Token
	current   *Token
	lookahead *Token
	pos       int
}

func (t tokenBuffer) Current() *Token {
	return t.current
}

func (t tokenBuffer) Lookahead() *Token {
	return t.lookahead
}

func NewTokenBuffer(t []*Token) *tokenBuffer {
	var lookahead *Token
	if len(t) != 0 {
		lookahead = t[0]
	}
	return &tokenBuffer{
		data:      t,
		current:   nil,
		lookahead: lookahead,
		pos:       -1,
	}
}
