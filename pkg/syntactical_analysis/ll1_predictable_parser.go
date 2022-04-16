package lexical_analysis

import (
	"context"
	"errors"
	"math-parser/pkg/entity"
	"math-parser/pkg/utils/logging"
)

func NewLL1PredictableParser(ctx context.Context) LL1PredictableParser {
	return &lL1PredictableParser{
		logging: ctx.Value("logger").(logging.Logger),
	}
}

type LL1PredictableParser interface {
	Parse([]*entity.Token) (entity.Ast, error)
}

type lL1PredictableParser struct {
	logging logging.Logger
	buffer  entity.TokenBuffer
}

func (l *lL1PredictableParser) Parse(t []*entity.Token) (entity.Ast, error) {
	l.buffer = entity.NewTokenBuffer(t)

	if root, err := l.makeExpression(); err == nil {
		return entity.NewAst(root), nil
	} else {
		return nil, err
	}

}

func (l *lL1PredictableParser) makeExpression() (entity.Node, error) {
	expr := entity.NewNode()
	currentToken := l.buffer.NextToken()

	switch currentToken.Tag {
	// EXPR -> TERM EXPR1
	case entity.OPERATOR_LEFT_BRACKET, entity.NUMBER:
		{

			term, err := l.makeTerm()
			if err != nil {
				return nil, errors.New("")
			}
			expr.AddChild(term)

			expr1, err := l.makeExpression1()
			if err != nil {
				return nil, errors.New("")
			}
			expr.AddChild(expr1)
		}
	default:
		{

		}
	}
	return expr, nil
}

func (l *lL1PredictableParser) makeFactor() (entity.Node, error) {
	return nil, nil
}

func (l *lL1PredictableParser) makeTerm() (entity.Node, error) {
	return nil, nil
}

func (l *lL1PredictableParser) makeExpression1() (entity.Node, error) {
	return nil, nil
}

func (l *lL1PredictableParser) makeTerm1() (entity.Node, error) {
	return nil, nil
}
