package math_parser

import (
	"context"
	"math-parser/pkg/compilation"
	"math-parser/pkg/lexical_analysis"
	"math-parser/pkg/preprocessing"
	"math-parser/pkg/utils/logging"
)

func main() {
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())

	preprocessor := preprocessing.NewPreprocessing()
	automata := lexical_analysis.NewAutomata(ctx)
	lexicalAnalyzer := lexical_analysis.NewLexicalAnalyzer(ctx, automata)

	compiler := compilation.NewCompiler(preprocessor, lexicalAnalyzer)
	compiler.Evaluate()
}
