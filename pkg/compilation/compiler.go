package compilation

import (
	"math-parser/pkg/lexical_analysis"
	"math-parser/pkg/preprocessing"
)

type compiler struct {
	preprocessor    preprocessing.Preprocessor
	lexicalAnalyzer lexical_analysis.LexicalAnalyzer
}

type Compiler interface {
	Evaluate()
}

func NewCompiler(pp preprocessing.Preprocessor, la lexical_analysis.LexicalAnalyzer) Compiler {
	return &compiler{
		preprocessor:    pp,
		lexicalAnalyzer: la,
	}
}

func (c *compiler) Evaluate() {
	return
}
