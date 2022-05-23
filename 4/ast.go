package main

import (
	"fmt"
	"log"
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

	Visit() interface{}
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

func (this *node) Visit() interface{} {
	return nil
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

/**
 * 语句
 * Statement 包括函数声明-FunctionDecl和函数调用-FunctionCall
 */

type Statement interface {
	ASTNode
}

type Expression interface {
	ASTNode
}

/**
 * 表达式语句
 * 就是在表达式后面加个分号
 */
type ExpressionStmt struct {
	ASTNode
}

func NewExpressionStmt(expr Expression) *ExpressionStmt {
	expr.SetKind(NodeKindExprStatement)
	return &ExpressionStmt{ASTNode: expr}
}

type FunctionCall struct {
	*node
	Name       string
	Parameters []string
	Definition *FunctionDecl
}

func NewFunctionCall(name string, parameters []string) *FunctionCall {
	return &FunctionCall{Name: name, Parameters: parameters}
}

func (this *FunctionCall) String() {
	log.Println("FunctionDecl Name:", this.Name, "Parameters:", this.Parameters, "Definition:", this.Definition)
}

func (this *FunctionCall) Kind() NodeKind {
	return NodeKindFunctionCall
}

type VariableDecl struct {
	*node
	name     string  //变量名称
	typ      string  //变量类型
	initExpr ASTNode //变量初始化所使用的表达式
}

func NewVariableDecl(name string, typ string, initExpr ASTNode) *VariableDecl {
	node := NewNode()
	node.SetKind(NodeKindVariable)
	return &VariableDecl{node: node, name: name, typ: typ, initExpr: initExpr}
}

type Variable struct {
	*node
	name string //变量名称
	decl *VariableDecl
}

func NewVariable(name string) *Variable {
	return &Variable{node: NewNode(), name: name}
}
func (this *Variable) Visit() interface{} {
	return nil
}

//*********FunctionDecl************/
type FunctionDecl struct {
	*node
	Name       string
	Parameters []string
	Body       *Block
}

func NewFunctionDecl(name string, params []string, body *Block) *FunctionDecl {
	return &FunctionDecl{Name: name, Parameters: params, Body: body}
}

func (this *FunctionDecl) String() {
	log.Println("FunctionDecl Name:", this.Name, "Parameters:", this.Parameters, "Body:", this.Body)
}

func (this *FunctionDecl) Kind() NodeKind {
	return NodeKindFunctionDecl
}

type Block struct {
	*node
	Stmts []Statement
}

func NewBlock(stmts []Statement) *Block {
	return &Block{node: NewNode(), Stmts: stmts}
}

func (this *Block) String() {
	for _, stmt := range this.Stmts {
		log.Println(stmt)
	}
}

func (this *Block) Kind() NodeKind {
	return NodeKindBlock
}

func (this *Block) Visit() interface{} {
	var ret interface{}
	for _, s := range this.Stmts {
		ret = s.Visit()
	}

	return ret
}

//func NewExpr(kind NodeKind, token *Token) *Expr {
//	expr := &Expr{NewNode()}
//	expr.SetLabel(token.Text)
//	expr.SetKind(kind)
//	expr.SetToken(token)
//	return expr
//}

type Prog struct {
	stmts []Statement
}

func NewProg(stmts []Statement) *Prog {
	prog := &Prog{stmts}
	return prog
}

func (this *Prog) String() {
	for _, stmt := range this.stmts {
		log.Println("Prog", stmt)
	}
}

func (this *Prog) Kind() NodeKind {
	return NodeKindProg
}

type ASTVisitor struct {
}

func NewASTVisitor() *ASTVisitor {
	return &ASTVisitor{}
}

func (this *ASTVisitor) Visit(node ASTNode) interface{} {
	return node.Visit()
}

//func (this *ASTVisitor) VisitProg(prog *Prog) interface{} {
//
//}
