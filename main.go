package math_parser

import (
	"math-parser/pkg/compilation"
	"math-parser/pkg/lexical_analysis"
	"math-parser/pkg/preprocessing"
)

func main() {
	preprocessor := preprocessing.NewPreprocessing()
	automata := lexical_analysis.NewAutomata()
	lexicalAnalyzer := lexical_analysis.NewLexicalAnalyzer(automata)

	compiler := compilation.NewCompiler(preprocessor, lexicalAnalyzer)
	compiler.Evaluate()
}
