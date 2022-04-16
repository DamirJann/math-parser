package preprocessing

import (
	"context"
	"math-parser/pkg/utils/logging"
	"strings"
)

const (
	SPACE string = " "
	EMPTY string = ""
)

func NewPreprocessing() Preprocessor {
	return preprocessor{
		logging: context.WithValue(context.Background(), "preprocessor", logging.NewBuiltinLogger()),
	}
}

type Preprocessor interface {
	tokenize(string) (output string)
}

type preprocessor struct {
	logging context.Context
}

func (preprocessor) tokenize(input string) (output string) {
	output = strings.ReplaceAll(input, SPACE, EMPTY)
	return output
}
