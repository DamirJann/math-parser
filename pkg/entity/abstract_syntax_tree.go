package entity

import (
	"fmt"
	"github.com/DamirJann/pretty-trie/pkg/drawing"
	"github.com/DamirJann/pretty-trie/pkg/entity"
)

type Ast interface {
	Evaluate() int
	Visualize() string
	Root() Node
}

type ast struct {
	root Node
}

type Node interface {
	Evaluate() (int, error)
	Token() *Token
	Label() string
	AddChildToEnd(...Node)
	AddChild(int, ...Node)
	AddChildToBegin(...Node)
	Delete(int)
	Child() []Node
	Replace(Node, int)
}

func NewAst(root Node) *ast {
	return &ast{
		root: root,
	}
}

func (n *node) AddChildToEnd(c ...Node) {
	n.child = append(n.child, c...)
}

func (n *node) AddChildToBegin(c ...Node) {
	n.child = append(c, n.child...)
}

func (n *node) AddChild(pos int, c ...Node) {
	n.child = append(append(n.child[:pos], c...), n.child[pos:]...)
}

func (n *node) Delete(pos int) {
	n.child = append(n.child[:pos], n.child[pos+1:]...)
}

func (n *node) Replace(new Node, pos int) {
	n.child[pos] = new
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
	//return a.root.Evaluate()
	return 0
}

func (a *ast) Root() Node {
	return a.root
}

func (n node) Evaluate() (res int, err error) {
	//for i, child := range n.child {
	//	switch child.Token().Tag {
	//	case OPERATOR_PLUS, OPERATOR_MINUS, OPERATOR_MULTIPLICATION, OPERATOR_DIVISION:
	//		{
	//			if i < 1 || i >= len(n.child)-1 {
	//				return 0, errors.New("not enough operands for bin.op")
	//			}
	//			lo, err := n.child[i-1].Evaluate()
	//			if err != nil {
	//				return 0, err
	//			}
	//			ro, err := n.child[i+1].Evaluate()
	//			if err != nil {
	//				return 0, err
	//			}
	//		}
	//	}
	//
	//}
	return 0, nil

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
