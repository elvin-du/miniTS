package main

import "log"

/**
 * 语句
 * Statement 包括函数声明-FunctionDecl和函数调用-FunctionCall
 */

type Statement interface {
	ASTNode
}

type FunctionCall struct {
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

//*********FunctionDecl************/
type FunctionDecl struct {
	Name       string
	Parameters []string
	Body       *FunctionBody
}

func NewFunctionDecl(name string, params []string, body *FunctionBody) *FunctionDecl {
	return &FunctionDecl{Name: name, Parameters: params, Body: body}
}

func (this *FunctionDecl) String() {
	log.Println("FunctionDecl Name:", this.Name, "Parameters:", this.Parameters, "Body:", this.Body)
}

func (this *FunctionDecl) Kind() NodeKind {
	return NodeKindFunctionDecl
}
