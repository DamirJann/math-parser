package lexical_analysis

import (
	"context"
	"gotest.tools/assert"
	"math-parser/pkg/entity"
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
			name:     "Happy flow. Parse substraction",
			scenario: happyFlowParseSubstraction,
		},
		{
			name:     "Happy flow. Parse expression with basic operations",
			scenario: happyFlowParseExprWithBasicOperations,
		},
		{
			name:     "Happy flow. Parse expression with brackets",
			scenario: happyFlowParseExprWithBrackets,
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

func happyFlowParseSubstraction(t *testing.T) {
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
