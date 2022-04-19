package main

import (
	"bufio"
	"context"
	"fmt"
	"math-parser/pkg/compilation"
	"math-parser/pkg/lexical_analysis"
	"math-parser/pkg/preprocessing"
	syntactical_analyzer "math-parser/pkg/syntactical_analysis"
	"math-parser/pkg/utils/logging"
	"os"
)

func main() {
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())

	preprocessor := preprocessing.NewPreprocessing()
	automata := lexical_analysis.NewAutomata()
	lexicalAnalyzer := lexical_analysis.NewLexicalAnalyzer(ctx, automata)
	syntacticalAnalyzer := syntactical_analyzer.NewLL1PredictableParser(ctx)
	compiler := compilation.NewCompiler(preprocessor, lexicalAnalyzer, syntacticalAnalyzer)

	in := bufio.NewReader(os.Stdin)
	expr, _, _ := in.ReadLine()
	res, err := compiler.Evaluate(string(expr))

	if err != nil {
		fmt.Printf("error: %s", err.Error())
	} else {
		fmt.Printf("Evaluation of `%s` is equal to %d", string(expr), res)
	}
}
