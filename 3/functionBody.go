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

func (this *FunctionBody) String() {
	for _, stmt := range this.Stmts {
		log.Println(stmt)
	}
}

func (this *FunctionBody) Kind() NodeKind {
	return NodeKindFunctionBody
}
