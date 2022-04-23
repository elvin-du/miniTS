package main

import "log"

/**
 * 语句
 * Statement 包括函数声明-FunctionDecl和函数调用-FunctionCall
 */

type Statement interface {
	AstNode
}

type FunctionCall struct {
	Name       string
	Parameters []string
	Definition *FunctionDecl
}

func NewFunctionCall(name string, parameters []string) *FunctionCall {
	return &FunctionCall{Name: name, Parameters: parameters}
}

func (this *FunctionCall) Dump(prefix string) {
	log.Println(prefix, "FunctionDecl", this.Name)
}

func (this *FunctionCall) Kind() NodeKind {
	return NodeKindFunctionCall
}

//*********FunctionDecl************/

type FunctionDecl struct {
	Name string
	Body *FunctionBody
}

func NewFunctionDecl(name string, body *FunctionBody) *FunctionDecl {
	return &FunctionDecl{Name: name, Body: body}
}

func (this *FunctionDecl) Dump(prefix string) {
	log.Println(prefix, "FunctionDecl", this.Name)
	this.Body.Dump(prefix)
}

func (this *FunctionDecl) Kind() NodeKind {
	return NodeKindFunctionDecl
}
