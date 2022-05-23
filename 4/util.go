package main

import "log"

func IsStatementNode(node ASTNode) bool {
	k := node.Kind()
	if k == NodeKindFunctionDecl || k == NodeKindFunctionCall {
		return true
	}

	return false
}

func IsFunctionCallNode(node ASTNode) bool {
	k := node.Kind()
	if k == NodeKindFunctionCall {
		return true
	}
	return false
}

func IsFunctionDeclNode(node ASTNode) bool {
	k := node.Kind()
	if k == NodeKindFunctionDecl {
		return true
	}
	return false
}

func IsFunctionBodyNode(node ASTNode) bool {
	k := node.Kind()
	if k == NodeKindBlock {
		return true
	}
	return false
}

type Stack struct {
	data []interface{}
	size int
	cap  int
}

func NewStack() *Stack {
	return &Stack{data: make([]interface{}, 1000), cap: 1000, size: 0}
}

func (this *Stack) Pop() interface{} {
	if this.size > 0 {
		this.size--
		return this.data[this.size]
	}

	return nil
}

func (this *Stack) Peek() interface{} {
	data := this.Pop()
	if nil == data {
		return nil
	}

	this.Push(data)
	return data
}

func (this *Stack) Push(i interface{}) {
	if this.size > this.cap {
		log.Fatalln("stack is full")
	}

	this.data[this.size] = i
	this.size++
}
