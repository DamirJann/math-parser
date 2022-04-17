package entity

import (
	"fmt"
	"github.com/DamirJann/pretty-trie/pkg/drawing"
	"github.com/DamirJann/pretty-trie/pkg/entity"
)

type Ast interface {
	Evaluate() int
	Visualize() string
}

type ast struct {
	root Node
}

type Node interface {
	Token() *Token
	Label() string
	AddChild(Node)
	Child() []Node
}

func NewAst(root Node) *ast {
	return &ast{
		root: root,
	}
}

func (n *node) AddChild(c Node) {
	n.child = append(n.child, c)
}

func (n node) Child() []Node {
	return n.child
}

func (n node) Label() string {
	return n.label
}
func (n node) Token() *Token {
	return n.token
}

type node struct {
	label string
	token *Token
	child []Node
}

func NewNonTerminalNode(l string) Node {
	return &node{
		label: l,
		token: nil,
		child: nil,
	}
}

func NewTerminalNode(t *Token) Node {
	return &node{
		token: t,
		child: nil,
	}
}

func NewEpsilonNode() Node {
	return &node{
		label: "eps",
		token: nil,
		child: nil,
	}
}

func (a *ast) Evaluate() int {
	return 0
}

func (a *ast) Visualize() string {
	idSeq := new(int)
	*idSeq = 0
	res, _ := drawing.Visualize(a.traverse(a.root, idSeq))
	return res
}

func (a *ast) traverse(n Node, idSeq *int) (node []entity.Node, edge []entity.Edge) {
	label := n.Label()
	if label == "" {
		label = fmt.Sprintf("%v", n.Token().Value)
	}
	node = append(node, entity.Node{
		Id:    *idSeq,
		Label: label,
	})
	for _, child := range n.Child() {
		edge = append(edge, entity.Edge{
			From: node[0].Id,
			To:   *idSeq + 1,
		})
		*idSeq++
		newNode, newEdge := a.traverse(child, idSeq)
		node = append(node, newNode...)
		edge = append(edge, newEdge...)
	}
	return node, edge
}
