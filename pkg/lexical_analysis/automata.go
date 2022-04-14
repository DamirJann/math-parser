package lexical_analysis

import (
	"bytes"
	"errors"
	"io"
)

type automata struct {
	input *bytes.Buffer
}

type Automata interface {
	extractToken(input *bytes.Buffer) (*token, error)
}

func NewAutomata() Automata {
	return &automata{}
}

func (a *automata) extractToken(input *bytes.Buffer) (*token, error) {
	a.input = input
	return a.s1("")
}

func (a *automata) s1(lexem string) (*token, error) {
	peek, err := a.Peek()

	if peek == EOF {
		return nil, io.EOF
	}

	if err != nil {
		return nil, err
	}

	if nextState := a.s1TransitTo(peek); nextState != nil {
		lexem += string(peek)
		return nextState(lexem)
	} else {
		return nil, errors.New("")
	}

}

func (a *automata) s2(lexem string) (*token, error) {

	if t, err := newNumberToken(lexem); err == nil {
		return t, nil
	} else {
		return nil, err
	}
}

func (a *automata) s3(lexem string) (*token, error) {
	peek, err := a.Peek()

	if err != nil {
		return nil, err
	}

	if nextState := a.s3TransitTo(peek); nextState != nil {
		return nextState(lexem + string(peek))
	} else {
		a.Unread()
		return a.s2(lexem)
	}

}

func (a *automata) s4(lexem string) (*token, error) {
	return newOperatorToken(lexem)
}

func (a *automata) s1TransitTo(lookahead byte) func(string) (*token, error) {
	return map[byte]func(string) (*token, error){
		ZERO: a.s2,

		PLUS:           a.s4,
		MINUS:          a.s4,
		MULTIPLICATION: a.s4,
		DIVISION:       a.s4,

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

func (a *automata) s3TransitTo(lookahead byte) func(string) (*token, error) {
	return map[byte]func(string) (*token, error){
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
}

func (a *automata) Peek() (byte, error) {
	b, err := a.input.ReadByte()
	if err == io.EOF {
		return EOF, nil
	}
	return b, err
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
