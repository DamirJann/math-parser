package lexical_analysis

import (
	"context"
	"gotest.tools/assert"
	"math-parser/pkg/utils/logging"
	"testing"
)

func TestLexicalAnalyzer_Tokenize(t *testing.T) {
	var tests = []struct {
		name     string
		scenario func(*testing.T)
	}{
		{
			name:     "Happy flow. Process with basic operations",
			scenario: happyFlowTokenizeWithBasicOperations,
		},
		{
			name:     "Happy flow. Process with all operations",
			scenario: happyFlowTokenizeWithAllOperations,
		},
		{
			name:     "Happy flow. Process empty expression",
			scenario: happyFlowTokenizeEmptyExpression,
		},
		{
			name:     "Happy flow. Process the only lexem",
			scenario: happyFlowTokenizeTheOnlyLexem,
		},
		{
			name:     "Happy flow. Process `0`",
			scenario: happyFlowTokenize0,
		},
		{
			name:     "Happy flow. Process 0123",
			scenario: happyFlowTokenize0123,
		},
		{
			name:     "Happy flow. Process expression with brackets",
			scenario: happyFlowTokenizeExpressionWithBrackets,
		},
	}

	t.Parallel()
	for _, test := range tests {
		t.Run(test.name, test.scenario)
	}
}

func happyFlowTokenizeWithBasicOperations(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	automata := NewAutomata()
	lexicalAnalyzer := NewLexicalAnalyzer(ctx, automata)
	expression := "123-555+0-3"

	// act
	ts, err := lexicalAnalyzer.Tokenize(expression)

	// assert
	assert.Equal(t, err, nil)
	assert.Equal(t, len(ts), 7)
	assert.Equal(t, ts[0].Value.(int), 123)
	assert.Equal(t, ts[1].Value.(string), "-")
	assert.Equal(t, ts[2].Value.(int), 555)
	assert.Equal(t, ts[3].Value.(string), "+")
	assert.Equal(t, ts[4].Value.(int), 0)
	assert.Equal(t, ts[5].Value.(string), "-")
	assert.Equal(t, ts[6].Value.(int), 3)
}

func happyFlowTokenizeWithAllOperations(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	automata := NewAutomata()
	lexicalAnalyzer := NewLexicalAnalyzer(ctx, automata)
	expression := "123*555/0+3-1"

	// act
	ts, err := lexicalAnalyzer.Tokenize(expression)

	// assert
	assert.Equal(t, err, nil)
	assert.Equal(t, len(ts), 9)
	assert.Equal(t, ts[0].Value.(int), 123)
	assert.Equal(t, ts[1].Value.(string), "*")
	assert.Equal(t, ts[2].Value.(int), 555)
	assert.Equal(t, ts[3].Value.(string), "/")
	assert.Equal(t, ts[4].Value.(int), 0)
	assert.Equal(t, ts[5].Value.(string), "+")
	assert.Equal(t, ts[6].Value.(int), 3)
	assert.Equal(t, ts[7].Value.(string), "-")
	assert.Equal(t, ts[8].Value.(int), 1)
}

func happyFlowTokenizeEmptyExpression(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())

	automata := NewAutomata()
	lexicalAnalyzer := NewLexicalAnalyzer(ctx, automata)

	// act
	ts, err := lexicalAnalyzer.Tokenize("")

	// assert
	assert.Equal(t, err, nil)
	assert.Equal(t, len(ts), 0)
}

func happyFlowTokenizeTheOnlyLexem(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())

	automata := NewAutomata()
	lexicalAnalyzer := NewLexicalAnalyzer(ctx, automata)

	// act
	ts, err := lexicalAnalyzer.Tokenize("12")

	// assert
	assert.Equal(t, err, nil)
	assert.Equal(t, len(ts), 1)
}

func happyFlowTokenize0(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())

	automata := NewAutomata()
	lexicalAnalyzer := NewLexicalAnalyzer(ctx, automata)

	// act
	ts, err := lexicalAnalyzer.Tokenize("0")

	// assert
	assert.Equal(t, err, nil)
	assert.Equal(t, len(ts), 1)
}

func happyFlowTokenize0123(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())

	automata := NewAutomata()
	lexicalAnalyzer := NewLexicalAnalyzer(ctx, automata)

	// act
	ts, err := lexicalAnalyzer.Tokenize("0123")

	// assert
	assert.Equal(t, err, nil)
	assert.Equal(t, len(ts), 2)
	assert.Equal(t, ts[0].Value.(int), 0)
	assert.Equal(t, ts[1].Value.(int), 123)
}

func happyFlowTokenizeExpressionWithBrackets(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())

	automata := NewAutomata()
	lexicalAnalyzer := NewLexicalAnalyzer(ctx, automata)

	// act
	ts, err := lexicalAnalyzer.Tokenize("(123+5)")

	// assert
	assert.Equal(t, err, nil)
	assert.Equal(t, len(ts), 5)
	assert.Equal(t, ts[0].Value.(string), "(")
	assert.Equal(t, ts[1].Value.(int), 123)
	assert.Equal(t, ts[2].Value.(string), "+")
	assert.Equal(t, ts[3].Value.(int), 5)
	assert.Equal(t, ts[4].Value.(string), ")")

}
