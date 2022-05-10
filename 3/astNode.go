package main

import (
	"fmt"
	"strings"
)

type ASTNode interface {
	//打印对象信息，通常用于缩进显示,int 表示几个空格
	Dump(int)
	Kind() NodeKind
	Children() []ASTNode
	Child(int) ASTNode
	Token() *Token
	Label() string

	SetKind(NodeKind)
	AddChild(ASTNode)
	SetToken(*Token)
	SetLabel(string)
}

type node struct {
	children []ASTNode
	kind     NodeKind
	token    *Token
	label    string
}

//test
var _ ASTNode = &node{}

func NewNode() *node {
	return &node{children: make([]ASTNode, 0)}
}
func (n *node) Dump(indent int) {
	fmt.Printf("%s%s\n", strings.Repeat("  ", indent*2), n.label)
	for _, child := range n.children {
		child.Dump(indent + 2)
	}
}

func (n *node) Kind() NodeKind {
	return n.kind
}
func (n *node) Children() []ASTNode {
	return n.children
}
func (n *node) Child(index int) ASTNode {
	if int(index) >= len(n.children) {
		return nil
	}
	return n.children[index]
}

func (n *node) Token() *Token {
	return n.token
}
func (n *node) Label() string {
	return n.label
}
func (n *node) AddChild(node ASTNode) {
	n.children = append(n.children, node)
}
func (n *node) SetKind(k NodeKind) {
	n.kind = k
}
func (n *node) SetToken(t *Token) {
	n.token = t
}
func (n *node) SetLabel(l string) {
	n.label = l
}
