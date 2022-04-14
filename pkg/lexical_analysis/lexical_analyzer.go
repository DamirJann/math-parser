package lexical_analysis

import (
	"bytes"
	"io"
)

func NewLexicalAnalyzer(automata Automata) LexicalAnalyzer {
	return &lexicalAnalyzer{
		automata: automata,
	}
}

type LexicalAnalyzer interface {
	Tokenize(input string) ([]token, error)
}

type lexicalAnalyzer struct {
	input    *bytes.Buffer
	automata Automata
}

func (la lexicalAnalyzer) Tokenize(input string) (output []token, err error) {
	la.input = bytes.NewBufferString(input)
	for {
		t, err := la.automata.extractToken(la.input)
		if err == io.EOF {
			break
		}
		if err != nil {
			return output, err
		}
		output = append(output, *t)
	}
	return output, nil
}
