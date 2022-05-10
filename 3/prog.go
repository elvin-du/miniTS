package main

import "log"

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
