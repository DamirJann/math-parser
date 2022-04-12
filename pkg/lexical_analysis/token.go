package lexical_analysis

type tokenName int

const (
	IDENTIFIER tokenName = iota
	NUMBER
	OPERATOR
)

type Token struct {
	name  tokenName
	value interface{}
}
