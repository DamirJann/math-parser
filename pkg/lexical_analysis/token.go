package lexical_analysis

import (
	"strconv"
)

type Tag int

const (
	IDENTIFIER Tag = iota
	NUMBER
	OPERATOR
)

type token struct {
	tag   Tag
	Value interface{}
}

func newNumberToken(lexem string) (*token, error) {
	if i, err := strconv.Atoi(lexem); err == nil {
		return &token{
			tag:   NUMBER,
			Value: i,
		}, nil
	} else {
		return nil, err
	}
}

func newOperatorToken(lexem string) (*token, error) {
	return &token{
		tag:   OPERATOR,
		Value: lexem,
	}, nil
}
