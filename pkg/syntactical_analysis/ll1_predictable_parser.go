package lexical_analysis

import (
	"context"
	"errors"
	"fmt"
	"math-parser/pkg/entity"
	"math-parser/pkg/utils/logging"
)

const (
	EXPR     = "EXPR"
	TERM     = "TERM"
	FACTOR   = "FACTOR"
	NUMBER   = "NUMBER"
	TERMINAL = "TERMINAL"
	EXPR1    = "EXPR1"
	TERM1    = "TERM1"
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
	l.bufferInit(t)

	if root, err := l.makeExpression(); err == nil {
		ast := entity.NewAst(root)
		l.logging.Debugf("computed ast: \n%v", ast.Visualize())

		l.simplify(ast)
		l.logging.Debugf("simplified ast: \n%v", ast.Visualize())

		return ast, nil
	} else {
		return nil, err
	}
}

func (l lL1PredictableParser) simplify(ast entity.Ast) {
	simplify(ast.Root())
}

func simplify(n entity.Node) {
	for i := len(n.Child()) - 1; i >= 0; i-- {
		simplify(n.Child()[i])
		if len(n.Child()[i].Child()) == 0 && n.Child()[i].Token() == nil {
			n.Delete(i)
		}

	}

	for i := len(n.Child()) - 1; i >= 0; i-- {
		if (n.Child()[i].Token() != nil) && (n.Child()[i].Token().Tag == entity.OPERATOR_LEFT_BRACKET || n.Child()[i].Token().Tag == entity.OPERATOR_RIGHT_BRACKET) {
			n.Delete(i)
		}
	}

	for i := len(n.Child()) - 1; i >= 0; i-- {
		if n.Child()[i].Token() == nil && len(n.Child()[i].Child()) == 1 {
			n.Replace(n.Child()[i].Child()[0], i)
		}
	}

	for i := len(n.Child()) - 1; i >= 0; i-- {
		switch n.Child()[i].Label() {
		case EXPR1:
			{

				//n.AddChild(n.Child()[i].Child()...)
				//n.Delete(i)
			}
		}
	}

}

func (l *lL1PredictableParser) bufferInit(t []*entity.Token) {
	l.buffer = entity.NewTokenBuffer(append(t, entity.NewEpsilonToken()))
}

func (l *lL1PredictableParser) makeExpression() (entity.Node, error) {
	expr := entity.NewNonTerminalNode(EXPR)

	switch l.buffer.Lookahead().Tag {
	case entity.OPERATOR_LEFT_BRACKET, entity.NUMBER:
		{
			term, err := l.makeTerm()
			if err != nil {
				return nil, err
			}
			expr.AddChild(term)

			expr1, err := l.makeExpression1()
			if err != nil {
				return nil, err
			}
			expr.AddChild(expr1)
			l.logging.Debugf("EXP -> TERM EXPR1 parsed")
		}
	default:
		{
			return nil, errors.New("error during production choosing for EXP")
		}
	}
	return expr, nil
}

func (l *lL1PredictableParser) makeFactor() (entity.Node, error) {
	factor := entity.NewNonTerminalNode(FACTOR)

	switch l.buffer.Lookahead().Tag {
	case entity.NUMBER:
		{
			l.buffer.NextToken()
			factor.AddChild(entity.NewTerminalNode(l.buffer.Current()))
			l.logging.Debugf("FACTOR -> number(%v) parsed", l.buffer.Current().Value)
		}
	case entity.OPERATOR_LEFT_BRACKET:
		{
			l.buffer.NextToken()
			factor.AddChild(entity.NewTerminalNode(l.buffer.Current()))

			expr, err := l.makeExpression()
			if err != nil {
				return nil, err
			}
			factor.AddChild(expr)

			if l.buffer.Lookahead().Tag != entity.OPERATOR_RIGHT_BRACKET {
				return nil, errors.New(fmt.Sprintf("expected ')', but found '%s'", l.buffer.Current().Value))
			}

			l.buffer.NextToken()
			factor.AddChild(entity.NewTerminalNode(l.buffer.Current()))

			l.logging.Debugf("FACTOR -> ( EXPR ) parsed")
		}
	default:
		{
			return nil, errors.New("error during production choosing for Factor")
		}
	}
	return factor, nil
}

func (l *lL1PredictableParser) makeTerm() (entity.Node, error) {
	term := entity.NewNonTerminalNode(TERM)

	switch l.buffer.Lookahead().Tag {
	case entity.OPERATOR_LEFT_BRACKET, entity.NUMBER:
		{
			factor, err := l.makeFactor()
			if err != nil {
				return nil, err
			}
			term.AddChild(factor)

			term1, err := l.makeTerm1()
			if err != nil {
				return nil, err
			}
			term.AddChild(term1)
			l.logging.Debugf("TERM -> FACTOR TERM1 parsed")
		}
	default:
		{
			return nil, errors.New("error during production choosing for TERM")
		}
	}
	return term, nil
}

func (l *lL1PredictableParser) makeExpression1() (entity.Node, error) {
	res := entity.NewNonTerminalNode(EXPR1)

	switch l.buffer.Lookahead().Tag {
	case entity.OPERATOR_PLUS:
		{
			l.buffer.NextToken()
			res.AddChild(entity.NewTerminalNode(l.buffer.Current()))
			term, err := l.makeTerm()
			if err != nil {
				return nil, err
			}
			res.AddChild(term)

			expr1, err := l.makeExpression1()
			if err != nil {
				return nil, err
			}
			res.AddChild(expr1)
			l.logging.Debugf("EXPR1 -> + TERM EXPR1 parsed", l.buffer.Current().Value)
		}
	case entity.OPERATOR_MINUS:
		{
			l.buffer.NextToken()
			res.AddChild(entity.NewTerminalNode(l.buffer.Current()))
			term, err := l.makeTerm()
			if err != nil {
				return nil, err
			}
			res.AddChild(term)

			expr1, err := l.makeExpression1()
			if err != nil {
				return nil, err
			}
			res.AddChild(expr1)
			l.logging.Debugf("EXPR1 -> - TERM EXPR1 parsed", l.buffer.Current().Value)
		}
	default:
		{
			res.AddChild(entity.NewEpsilonNode())
			l.logging.Debugf("EXPRESSION1 -> EPSILON parsed", l.buffer.Current().Value)
		}
	}
	return res, nil
}

func (l *lL1PredictableParser) makeTerm1() (entity.Node, error) {
	res := entity.NewNonTerminalNode(TERM1)

	switch l.buffer.Lookahead().Tag {
	case entity.OPERATOR_MULTIPLICATION:
		{
			l.buffer.NextToken()
			res.AddChild(entity.NewTerminalNode(l.buffer.Current()))
			factor, err := l.makeFactor()
			if err != nil {
				return nil, err
			}
			res.AddChild(factor)

			term1, err := l.makeTerm1()
			if err != nil {
				return nil, err
			}
			res.AddChild(term1)
			l.logging.Debugf("TERM1 -> * TERM TERM1 parsed", l.buffer.Current().Value)
		}
	case entity.OPERATOR_DIVISION:
		{
			l.buffer.NextToken()
			res.AddChild(entity.NewTerminalNode(l.buffer.Current()))
			factor, err := l.makeFactor()
			if err != nil {
				return nil, err
			}
			res.AddChild(factor)

			term1, err := l.makeTerm1()
			if err != nil {
				return nil, err
			}
			res.AddChild(term1)
			l.logging.Debugf("TERM1 -> / TERM TERM1 parsed", l.buffer.Current().Value)
		}
	default:
		{
			res.AddChild(entity.NewEpsilonNode())
			l.logging.Debugf("TERM1 -> EPSILON parsed", l.buffer.Current().Value)
		}
	}
	return res, nil
}
