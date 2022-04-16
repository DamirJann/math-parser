package lexical_analysis

import (
	"bytes"
	"context"
	"io"
	"math-parser/pkg/entity"
	"math-parser/pkg/utils/logging"
)

func NewLexicalAnalyzer(ctx context.Context, automata Automata) LexicalAnalyzer {
	return &lexicalAnalyzer{
		automata: automata,
		logging:  ctx.Value("logger").(logging.Logger),
	}
}

type LexicalAnalyzer interface {
	Tokenize(input string) ([]entity.Token, error)
}

type lexicalAnalyzer struct {
	input    *bytes.Buffer
	automata Automata
	logging  logging.Logger
}

func (la *lexicalAnalyzer) Tokenize(input string) (output []entity.Token, err error) {
	la.initBuffer(input)
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
	la.logging.Debugf("tokens len=%d: %v", len(output), output)
	return output, nil
}

func (la *lexicalAnalyzer) initBuffer(input string) {
	la.input = bytes.NewBufferString(input)
	la.input.WriteByte(EOF)
}
