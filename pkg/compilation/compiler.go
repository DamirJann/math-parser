package compilation

import (
	"math-parser/pkg/lexical_analysis"
	"math-parser/pkg/preprocessing"
	syntactical_analyzer "math-parser/pkg/syntactical_analysis"
)

type compiler struct {
	preprocessor        preprocessing.Preprocessor
	lexicalAnalyzer     lexical_analysis.LexicalAnalyzer
	syntacticalAnalyzer syntactical_analyzer.LL1PredictableParser
}

type Compiler interface {
	Evaluate(input string) (int, error)
}

func NewCompiler(pp preprocessing.Preprocessor, la lexical_analysis.LexicalAnalyzer, sa syntactical_analyzer.LL1PredictableParser) Compiler {
	return &compiler{
		preprocessor:        pp,
		lexicalAnalyzer:     la,
		syntacticalAnalyzer: sa,
	}
}

func (c *compiler) Evaluate(input string) (int, error) {
	input = c.preprocessor.Process(input)
	ast, err := c.lexicalAnalyzer.Tokenize(input)
	if err != nil {
		return 0, err
	}
	res, err := c.syntacticalAnalyzer.Evaluate(ast)
	if err != nil {
		return 0, err
	}
	return res, nil
}
