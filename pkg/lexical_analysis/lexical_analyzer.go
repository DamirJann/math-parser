package lexical_analysis

import (
	"bytes"
	"errors"
)

func NewLexicalAnalyzer(input bytes.Buffer) LexicalAnalyzer {
	return &lexicalAnalyzer{
		input: input,
	}
}

type LexicalAnalyzer interface {
	Tokenize() []*Token
	Peek() (string, error)
	Lookahead() (string, error)
	Skip() error
}

type lexicalAnalyzer struct {
	input bytes.Buffer
}

func (la *lexicalAnalyzer) Peek() (string, error) {
	b, err := la.input.ReadByte()
	return string(b), err
}

func (la *lexicalAnalyzer) Lookahead() (string, error) {
	b, err := la.input.ReadByte()
	if err == nil {
		err = la.input.UnreadByte()
	}
	return string(b), err
}

func (la *lexicalAnalyzer) Skip() error {
	_, err := la.input.ReadByte()
	return err
}

func (la lexicalAnalyzer) Tokenize() (output []*Token) {
	for {
		t, err := la.runAutomata()
		if err != nil {
			break
		}
		output = append(output, t)
	}
	return output
}

func (la lexicalAnalyzer) runAutomata() (*Token, error) {
	return la.state1()
}

func (la lexicalAnalyzer) state1() (*Token, error) {
	smb, err := la.Peek()
	if err != nil {
		return nil, err
	}

	switch smb {
	case ZERO:
		{
			return la.state2()
		}
	case ONE, TWO, THREE, FOUR, FIVE, SIX, SEVEN, EIGHT, NINE:
		{
			return la.state3(smb)
		}
	case PLUS, MINUS, MULTIPLICATION, DIVISION:
		{
			return la.state3(smb)
		}
	default:
		{
			return nil, errors.New("")
		}
	}

}

func (la lexicalAnalyzer) state2() (*Token, error) {
	return &Token{name: NUMBER, value: ZERO}, nil
}

func (la lexicalAnalyzer) state4(op string) (*Token, error) {
	return &Token{name: OPERATOR, value: op}, nil
}

func (la lexicalAnalyzer) state3(storage string) (*Token, error) {
	lookahead, err := la.Lookahead()
	if err != nil {
		return nil, err
	}

	switch lookahead {
	case ZERO, ONE, TWO, THREE, FOUR, FIVE, SIX, SEVEN, EIGHT, NINE:
		{
			if err := la.Skip(); err != nil {
				return nil, err
			}
			return la.state3(storage + lookahead)
		}
	default:
		{
			return &Token{name: NUMBER, value: storage}, nil
		}
	}
}
