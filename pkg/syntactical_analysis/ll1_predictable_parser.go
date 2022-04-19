package syntactical_analyzer

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
	Evaluate(t []*entity.Token) (int, error)
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

func (l *lL1PredictableParser) Evaluate(t []*entity.Token) (int, error) {
	ast, err := l.Parse(t)
	if err != nil {
		return 0, err
	}

	postfix, err := toPostfix(ast.Root())
	if err != nil {
		return 0, err
	}

	l.logging.Debug("postfix notation: %v", postfix)
	return evaluatePostfix(postfix)
}

func evaluatePostfix(ts []entity.Token) (int, error) {
	stack := entity.NewStack()

	for _, t := range ts {
		if t.Tag == entity.NUMBER {
			stack.Push(t)
			continue
		}

		r := stack.Pop()
		l := stack.Pop()
		switch t.Tag {
		case entity.OPERATOR_PLUS:
			{
				stack.Push(entity.Token{entity.NUMBER, l.Value.(int) + r.Value.(int)})
			}
		case entity.OPERATOR_MINUS:
			{
				stack.Push(entity.Token{entity.NUMBER, l.Value.(int) - r.Value.(int)})
			}
		case entity.OPERATOR_MULTIPLICATION:
			{
				stack.Push(entity.Token{entity.NUMBER, l.Value.(int) * r.Value.(int)})
			}
		case entity.OPERATOR_DIVISION:
			{
				stack.Push(entity.Token{entity.NUMBER, l.Value.(int) / r.Value.(int)})
			}
		}
	}
	return stack.Pop().Value.(int), nil
}

func toPostfix(n entity.Node) (res []entity.Token, err error) {
	if n.Token() != nil && n.Token().Tag == entity.NUMBER {
		return []entity.Token{*n.Token()}, nil
	}

	first, err := toPostfix(n.Child()[0])
	if err != nil {
		return nil, err
	}
	res = append(res, first...)
	var lastOp *entity.Token
	for i := 1; i < len(n.Child()); i++ {

		child := n.Child()[i]

		if child.Token() == nil || child.Token().Tag == entity.NUMBER {
			r, err := toPostfix(child)
			if err != nil {
				return nil, err
			}
			res = append(res, r...)
			res = append(res, *lastOp)

		} else {

			switch child.Token().Tag {
			case entity.OPERATOR_PLUS, entity.OPERATOR_MINUS, entity.OPERATOR_DIVISION, entity.OPERATOR_MULTIPLICATION:
				{
					lastOp = child.Token()
				}
			default:
				return nil, errors.New(fmt.Sprintf("unknown token:%v", *child.Token()))
			}
		}

	}
	return res, nil
}

func evaluate(n entity.Node) (int, error) {
	if n.Token() != nil {
		switch n.Token().Tag {
		case entity.NUMBER:
			{
				return n.Token().Value.(int), nil
			}
		}
		return 0, errors.New("trying evaluate not number")
	} else {
		lo, err := evaluate(n.Child()[0])
		if err != nil {
			return 0, err
		}
		ro, err := evaluate(n.Child()[2])
		if err != nil {
			return 0, err
		}
		switch n.Child()[1].Token().Tag {
		case entity.OPERATOR_PLUS:
			{
				return lo + ro, nil
			}
		case entity.OPERATOR_DIVISION:
			{
				return lo / ro, nil
			}
		case entity.OPERATOR_MULTIPLICATION:
			{
				return lo * ro, nil
			}
		case entity.OPERATOR_MINUS:
			{
				return lo - ro, nil
			}
		default:
			return 0, errors.New("unknown operation")
		}
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
	//
	for i := len(n.Child()) - 1; i >= 0; i-- {
		if n.Child()[i].Token() == nil && len(n.Child()[i].Child()) == 1 {
			n.Replace(n.Child()[i].Child()[0], i)
		}
	}

	for i := len(n.Child()) - 1; i >= 0; i-- {
		if n.Child()[i].Label() == EXPR1 || n.Child()[i].Label() == TERM1 {
			n.AddChild(i+1, n.Child()[i].Child()...)
			n.Delete(i)
		}
	}

	//for i := len(n.Child()) - 1; i >= 0; i-- {
	//	if n.Label() == EXPR && (n.Child()[i].Label() == EXPR1 || n.Child()[i].Label() == TERM1) {
	//		n.Child()[i+1].AddChildToBegin(n.Child()[i])
	//		n.Delete(i)
	//		break
	//	}
	//}
	//
	//if len(n.Child()) == 1 && n.Token() == nil && n.Child()[0].Token() == nil {
	//	n.AddChildToEnd(n.Child()[0].Child()...)
	//	n.Delete(0)
	//}

	//for i := len(n.Child()) - 1; i >= 0; i-- {
	//	if n.Child()[i].Label() != EXPR || len(n.Child()[i].Child()) == 0 {
	//		continue
	//	}
	//	switch n.Child()[i].Child()[0].Token().Tag {
	//	case entity.OPERATOR_PLUS, entity.OPERATOR_MINUS:
	//		{
	//			op := n.Child()[i].Child()[0]
	//			n.Child()[i].Delete(0)
	//			``
	//		}
	//
	//	}
	//}

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
			expr.AddChildToEnd(term)

			expr1, err := l.makeExpression1()
			if err != nil {
				return nil, err
			}
			expr.AddChildToEnd(expr1)
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
			factor.AddChildToEnd(entity.NewTerminalNode(l.buffer.Current()))
			l.logging.Debugf("FACTOR -> number(%v) parsed", l.buffer.Current().Value)
		}
	case entity.OPERATOR_LEFT_BRACKET:
		{
			l.buffer.NextToken()
			factor.AddChildToEnd(entity.NewTerminalNode(l.buffer.Current()))

			expr, err := l.makeExpression()
			if err != nil {
				return nil, err
			}
			factor.AddChildToEnd(expr)

			if l.buffer.Lookahead().Tag != entity.OPERATOR_RIGHT_BRACKET {
				return nil, errors.New(fmt.Sprintf("expected ')', but found '%s'", l.buffer.Current().Value))
			}

			l.buffer.NextToken()
			factor.AddChildToEnd(entity.NewTerminalNode(l.buffer.Current()))

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
			term.AddChildToEnd(factor)

			term1, err := l.makeTerm1()
			if err != nil {
				return nil, err
			}
			term.AddChildToEnd(term1)
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
			res.AddChildToEnd(entity.NewTerminalNode(l.buffer.Current()))
			term, err := l.makeTerm()
			if err != nil {
				return nil, err
			}
			res.AddChildToEnd(term)

			expr1, err := l.makeExpression1()
			if err != nil {
				return nil, err
			}
			res.AddChildToEnd(expr1)
			l.logging.Debugf("EXPR1 -> + TERM EXPR1 parsed", l.buffer.Current().Value)
		}
	case entity.OPERATOR_MINUS:
		{
			l.buffer.NextToken()
			res.AddChildToEnd(entity.NewTerminalNode(l.buffer.Current()))
			term, err := l.makeTerm()
			if err != nil {
				return nil, err
			}
			res.AddChildToEnd(term)

			expr1, err := l.makeExpression1()
			if err != nil {
				return nil, err
			}
			res.AddChildToEnd(expr1)
			l.logging.Debugf("EXPR1 -> - TERM EXPR1 parsed", l.buffer.Current().Value)
		}
	default:
		{
			res.AddChildToEnd(entity.NewEpsilonNode())
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
			res.AddChildToEnd(entity.NewTerminalNode(l.buffer.Current()))
			factor, err := l.makeFactor()
			if err != nil {
				return nil, err
			}
			res.AddChildToEnd(factor)

			term1, err := l.makeTerm1()
			if err != nil {
				return nil, err
			}
			res.AddChildToEnd(term1)
			l.logging.Debugf("TERM1 -> * TERM TERM1 parsed", l.buffer.Current().Value)
		}
	case entity.OPERATOR_DIVISION:
		{
			l.buffer.NextToken()
			res.AddChildToEnd(entity.NewTerminalNode(l.buffer.Current()))
			factor, err := l.makeFactor()
			if err != nil {
				return nil, err
			}
			res.AddChildToEnd(factor)

			term1, err := l.makeTerm1()
			if err != nil {
				return nil, err
			}
			res.AddChildToEnd(term1)
			l.logging.Debugf("TERM1 -> / TERM TERM1 parsed", l.buffer.Current().Value)
		}
	default:
		{
			res.AddChildToEnd(entity.NewEpsilonNode())
			l.logging.Debugf("TERM1 -> EPSILON parsed", l.buffer.Current().Value)
		}
	}
	return res, nil
}
