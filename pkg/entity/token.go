package entity

import (
	"strconv"
)

type Tag int

const (
	IDENTIFIER Tag = iota
	NUMBER
	OPERATOR_PLUS
	OPERATOR_MINUS
	OPERATOR_MULTIPLICATION
	OPERATOR_DIVISION
	OPERATOR_LEFT_BRACKET
	OPERATOR_RIGHT_BRACKET

	EPSILON
)

type Token struct {
	Tag   Tag
	Value interface{}
}

func NewNumberToken(lexem string) (*Token, error) {
	if i, err := strconv.Atoi(lexem); err == nil {
		return &Token{
			Tag:   NUMBER,
			Value: i,
		}, nil
	} else {
		return nil, err
	}
}

func NewOperatorToken(lexem string) (*Token, error) {
	tokenTag := map[string]Tag{
		"+": OPERATOR_PLUS,
		"-": OPERATOR_MINUS,
		"*": OPERATOR_MULTIPLICATION,
		"/": OPERATOR_DIVISION,
		"(": OPERATOR_LEFT_BRACKET,
		")": OPERATOR_RIGHT_BRACKET,
	}
	return &Token{
		Tag:   tokenTag[lexem],
		Value: lexem,
	}, nil
}

func NewEpsilonToken() *Token {
	return &Token{
		Tag:   EPSILON,
		Value: nil,
	}
}
