package lexical_analysis

import (
	"bytes"
	"errors"
	"io"
	"math-parser/pkg/entity"
)

type automata struct {
	input *bytes.Buffer
	lexem string
}

type Automata interface {
	extractToken(input *bytes.Buffer) (*entity.Token, error)
}

func NewAutomata() Automata {
	return &automata{}
}

func (a *automata) Peek() (byte, error) {
	return a.input.ReadByte()
}

func (a *automata) Lookahead() (byte, error) {
	b, err := a.input.ReadByte()
	if err == nil {
		err = a.input.UnreadByte()
	}
	if err == io.EOF {
		return EOF, nil
	}
	return b, err
}

func (a *automata) Unread() error {
	return a.input.UnreadByte()
}

func (a *automata) extractToken(input *bytes.Buffer) (*entity.Token, error) {
	a.input = input
	a.lexem = ""

	lookahead, err := a.Lookahead()
	if lookahead == EOF {
		return nil, io.EOF
	}
	if err != nil {
		return nil, err
	}

	return a.s1()
}

func (a *automata) s1() (*entity.Token, error) {
	peek, err := a.Peek()

	if err != nil {
		return nil, err
	}

	if nextState := a.s1TransitTo(peek); nextState != nil {
		a.lexem += string(peek)
		return nextState()
	} else {
		return nil, errors.New("[LOG][Automata]: Error in S1 state")
	}

}

func (a *automata) s2() (*entity.Token, error) {

	if t, err := entity.NewNumberToken(a.lexem); err == nil {
		return t, nil
	} else {
		return nil, err
	}
}

func (a *automata) s3() (*entity.Token, error) {
	peek, err := a.Peek()

	if err != nil {
		return nil, err
	}

	a.lexem += string(peek)
	nextState := a.s3TransitTo(peek)
	return nextState()
}

func (a *automata) s4() (*entity.Token, error) {
	return entity.NewOperatorToken(a.lexem)
}

func (a *automata) s5() (*entity.Token, error) {
	if err := a.Unread(); err != nil {
		return nil, err
	}
	return entity.NewNumberToken(a.lexem[0 : len(a.lexem)-1])

}

func (a *automata) s1TransitTo(lookahead byte) func() (*entity.Token, error) {
	return map[byte]func() (*entity.Token, error){
		ZERO: a.s2,

		PLUS:           a.s4,
		MINUS:          a.s4,
		MULTIPLICATION: a.s4,
		DIVISION:       a.s4,
		LEFT_BRACKET:   a.s4,
		RIGHT_BRACKET:  a.s4,

		ONE:   a.s3,
		TWO:   a.s3,
		THREE: a.s3,
		FOUR:  a.s3,
		FIVE:  a.s3,
		SIX:   a.s3,
		SEVEN: a.s3,
		EIGHT: a.s3,
		NINE:  a.s3,
	}[lookahead]
}

func (a *automata) s3TransitTo(lookahead byte) func() (*entity.Token, error) {
	next, ok := map[byte]func() (*entity.Token, error){
		ZERO:  a.s3,
		ONE:   a.s3,
		TWO:   a.s3,
		THREE: a.s3,
		FOUR:  a.s3,
		FIVE:  a.s3,
		SIX:   a.s3,
		SEVEN: a.s3,
		EIGHT: a.s3,
		NINE:  a.s3,
	}[lookahead]
	if ok {
		return next
	}
	return a.s5
}
