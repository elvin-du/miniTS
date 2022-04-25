package main

import "log"

type Prog struct {
	stmts []Statement
}

func NewProg(stmts []Statement) *Prog {
	prog := &Prog{stmts}
	return prog
}

func (this *Prog) Dump(prefix string) {
	log.Println(prefix, "Prog") //todo
}

func (this *Prog) Kind() NodeKind {
	return NodeKindProg
}
