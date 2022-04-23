package main

import (
	"log"
)

type FunctionBody struct {
	Stmts []Statement
}

func NewFunctionBody(stmts []Statement) *FunctionBody {
	return &FunctionBody{Stmts: stmts}
}

func (this FunctionBody) Dump(prefix string) {
	log.Println(prefix, "") //todo
}

func (this FunctionBody) Kind() NodeKind {
	return NodeKindFunctionBody
}
