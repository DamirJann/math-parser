package preprocessing

import "strings"

const (
	SPACE string = " "
	EMPTY string = ""
)

func NewPreprocessing() Preprocessor {
	return preprocessor{}
}

type Preprocessor interface {
	tokenize(string) (output string)
}

type preprocessor struct{}

func (preprocessor) tokenize(input string) (output string) {
	output = strings.ReplaceAll(input, SPACE, EMPTY)
	return output
}
