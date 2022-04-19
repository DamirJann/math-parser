package preprocessing

import (
	"gotest.tools/assert"
	"testing"
)

func TestPreprocessor_Process(t *testing.T) {
	var tests = []struct {
		Name     string
		Scenario func(*testing.T)
	}{
		{
			Name:     "Happy flow. Get rid of space",
			Scenario: happyFlowGetRidOfSpace,
		},
	}

	for _, test := range tests {
		t.Parallel()
		t.Run(test.Name, test.Scenario)
	}
}

func happyFlowGetRidOfSpace(t *testing.T) {
	// arrange
	preprocessor := NewPreprocessing()
	input := "   4+5-1 + 10 - var1   - 1   "

	// act
	output := preprocessor.Process(input)

	// assert
	assert.Equal(t, output, "4+5-1+10-var1-1")
}
