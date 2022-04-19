package syntactical_analyzer

import (
	"context"
	"gotest.tools/assert"
	"math-parser/pkg/entity"
	"math-parser/pkg/lexical_analysis"
	"math-parser/pkg/utils/logging"
	"testing"
)

func TestLexicalAnalyzer_Tokenize(t *testing.T) {
	var tests = []struct {
		name     string
		scenario func(*testing.T)
	}{
		{
			name:     "Happy flow. Parse addition",
			scenario: happyFlowParseAddition,
		},
		{
			name:     "Happy flow. Parse subtraction",
			scenario: happyFlowParseSubtraction,
		},
		{
			name:     "Happy flow. Parse expression with basic operations",
			scenario: happyFlowParseExprWithBasicOperations,
		},
		{
			name:     "Happy flow. Parse expression with brackets",
			scenario: happyFlowParseExprWithBrackets,
		},
		{
			name:     "Happy flow. Parse complicated expression",
			scenario: happyFlowParseComplicatedExpression,
		},
	}

	t.Parallel()
	for _, test := range tests {
		t.Run(test.name, test.scenario)
	}
}

func happyFlowParseAddition(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	parser := NewLL1PredictableParser(ctx)
	tk := []*entity.Token{
		{Tag: entity.NUMBER, Value: 1},
		{Tag: entity.OPERATOR_PLUS, Value: "+"},
		{Tag: entity.NUMBER, Value: 10},
	}
	// act
	_, err := parser.Parse(tk)
	// assert
	assert.Equal(t, err, nil)
}

func happyFlowParseSubtraction(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	parser := NewLL1PredictableParser(ctx)
	tk := []*entity.Token{
		{Tag: entity.NUMBER, Value: 1},
		{Tag: entity.OPERATOR_MINUS, Value: "-"},
		{Tag: entity.NUMBER, Value: 10},
	}
	// act
	_, err := parser.Parse(tk)

	// assert
	assert.Equal(t, err, nil)
}

func happyFlowParseExprWithBasicOperations(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	parser := NewLL1PredictableParser(ctx)
	tk := []*entity.Token{
		{Tag: entity.NUMBER, Value: 1},
		{Tag: entity.OPERATOR_MINUS, Value: "-"},
		{Tag: entity.NUMBER, Value: 10},
		{Tag: entity.OPERATOR_PLUS, Value: "+"},
		{Tag: entity.NUMBER, Value: 30},
		{Tag: entity.OPERATOR_DIVISION, Value: "/"},
		{Tag: entity.NUMBER, Value: 50},
		{Tag: entity.OPERATOR_MULTIPLICATION, Value: "*"},
		{Tag: entity.NUMBER, Value: 60},
	}
	// act
	_, err := parser.Parse(tk)

	// assert
	assert.Equal(t, err, nil)
}

func happyFlowParseExprWithBrackets(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	parser := NewLL1PredictableParser(ctx)
	tk := []*entity.Token{
		{Tag: entity.NUMBER, Value: 1},
		{Tag: entity.OPERATOR_MINUS, Value: "-"},
		{Tag: entity.OPERATOR_LEFT_BRACKET, Value: "("},
		{Tag: entity.NUMBER, Value: 10},
		{Tag: entity.OPERATOR_PLUS, Value: "+"},
		{Tag: entity.NUMBER, Value: 30},
		{Tag: entity.OPERATOR_RIGHT_BRACKET, Value: ")"},
		{Tag: entity.OPERATOR_DIVISION, Value: "/"},
		{Tag: entity.NUMBER, Value: 50},
		{Tag: entity.OPERATOR_MULTIPLICATION, Value: "*"},
		{Tag: entity.OPERATOR_LEFT_BRACKET, Value: "("},
		{Tag: entity.NUMBER, Value: 60},
		{Tag: entity.OPERATOR_RIGHT_BRACKET, Value: ")"},
	}
	// act
	_, err := parser.Parse(tk)

	// assert
	assert.Equal(t, err, nil)
}

func happyFlowParseComplicatedExpression(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	automata := lexical_analysis.NewAutomata()
	analyzer := lexical_analysis.NewLexicalAnalyzer(ctx, automata)
	parser := NewLL1PredictableParser(ctx)

	// act
	tk, _ := analyzer.Tokenize("(4+23/(42-(4-6+2))-0-1+9)")
	_, err := parser.Parse(tk)

	// assert
	assert.Equal(t, err, nil)
}

func TestLexicalAnalyzer_Evaluate(t *testing.T) {
	var tests = []struct {
		name     string
		scenario func(*testing.T)
	}{
		{
			name:     "Happy flow. Evaluate addition",
			scenario: happyFlowEvaluateAddition,
		},
		{
			name:     "Happy flow. Evaluate substraction",
			scenario: happyFlowEvaluateSubstraction,
		},
		{
			name:     "Happy flow. Evaluate one number",
			scenario: happyFlowEvaluateOneNumber,
		},
		{
			name:     "Happy flow. Evaluate expression with basic operations",
			scenario: happyFlowEvaluateExprWithBasicOperations,
		},
		{
			name:     "Happy flow. Evaluate expression with brackets",
			scenario: happyFlowEvaluateExprWithBrackets,
		},
		{
			name:     "Happy flow. Evaluate complicated expression",
			scenario: happyFlowEvaluateComplicatedExpression,
		},
	}

	t.Parallel()
	for _, test := range tests {
		t.Run(test.name, test.scenario)
	}
}

func happyFlowEvaluateAddition(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	parser := NewLL1PredictableParser(ctx)
	tk := []*entity.Token{
		{Tag: entity.NUMBER, Value: 1},
		{Tag: entity.OPERATOR_PLUS, Value: "+"},
		{Tag: entity.NUMBER, Value: 10},
	}
	// act
	res, err := parser.Evaluate(tk)
	// assert
	assert.Equal(t, err, nil)
	assert.Equal(t, res, 11)
}

func happyFlowEvaluateSubstraction(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	parser := NewLL1PredictableParser(ctx)
	tk := []*entity.Token{
		{Tag: entity.NUMBER, Value: 1},
		{Tag: entity.OPERATOR_MINUS, Value: "-"},
		{Tag: entity.NUMBER, Value: 10},
	}
	// act
	res, err := parser.Evaluate(tk)

	// assert
	assert.Equal(t, err, nil)
	assert.Equal(t, res, -9)
}

func happyFlowEvaluateExprWithBasicOperations(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	parser := NewLL1PredictableParser(ctx)
	tk := []*entity.Token{
		{Tag: entity.NUMBER, Value: 1},
		{Tag: entity.OPERATOR_MINUS, Value: "-"},
		{Tag: entity.NUMBER, Value: 10},
		{Tag: entity.OPERATOR_PLUS, Value: "+"},
		{Tag: entity.NUMBER, Value: 100},
		{Tag: entity.OPERATOR_DIVISION, Value: "/"},
		{Tag: entity.NUMBER, Value: 50},
		{Tag: entity.OPERATOR_MULTIPLICATION, Value: "*"},
		{Tag: entity.NUMBER, Value: 60},
	}
	// act
	res, err := parser.Evaluate(tk)

	// assert
	assert.Equal(t, err, nil)
	assert.Equal(t, res, 111)
}

func happyFlowEvaluateExprWithBrackets(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	automata := lexical_analysis.NewAutomata()
	analyzer := lexical_analysis.NewLexicalAnalyzer(ctx, automata)
	parser := NewLL1PredictableParser(ctx)

	// act
	tk, _ := analyzer.Tokenize("1-(10+30)/10*(60)")
	res, err := parser.Evaluate(tk)

	// assert
	assert.Equal(t, err, nil)
	assert.Equal(t, res, -239)
}

func happyFlowEvaluateComplicatedExpression(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	automata := lexical_analysis.NewAutomata()
	analyzer := lexical_analysis.NewLexicalAnalyzer(ctx, automata)
	parser := NewLL1PredictableParser(ctx)

	// act
	tk, _ := analyzer.Tokenize("(4+84/(42-(4-6+2))-0-1+9)")
	res, err := parser.Evaluate(tk)

	// assert
	assert.Equal(t, err, nil)
	assert.Equal(t, res, 14)
}

func happyFlowEvaluateOneNumber(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	automata := lexical_analysis.NewAutomata()
	analyzer := lexical_analysis.NewLexicalAnalyzer(ctx, automata)
	parser := NewLL1PredictableParser(ctx)

	// act
	tk, _ := analyzer.Tokenize("1000")
	res, err := parser.Evaluate(tk)

	// assert
	assert.Equal(t, err, nil)
	assert.Equal(t, res, 1000)
}
