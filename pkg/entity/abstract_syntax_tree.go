package entity

type Ast interface {
	Evaluate() int
}

type ast struct {
	root Node
}

type Node interface {
	AddChild(Node)
}

func NewAst(root Node) *ast {
	return &ast{
		root: root,
	}
}

func (n *node) AddChild(c Node) {
	n.child = append(n.child, c)
}

type node struct {
	Value Token
	child []Node
}

func NewNode() *node {
	return &node{
		Value: Token{},
		child: nil,
	}
}

func (a *ast) Evaluate() int {
	return 0
}
