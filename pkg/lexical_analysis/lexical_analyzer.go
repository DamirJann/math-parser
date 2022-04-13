package lexical_analysis

import "io"

func NewLexicalAnalyzer(input io.Reader) LexicalAnalyzer {
	return lexicalAnalyzer{
		lookahead: input,
	}
}

type LexicalAnalyzer interface {
	tokenize() []Token
}

type lexicalAnalyzer struct {
	lookahead io.Reader
}

func (lexicalAnalyzer) tokenize() []Token {
	return []Token{}
}
