package main

import "log"

type RefResolver struct {
	VisitorAst
}

func NewRefResolver() *RefResolver {
	return &RefResolver{}
}

func (this RefResolver) resolveFunctionCall(prog *Prog, name string) {
	for _, stmt := range prog.stmts {
		if IsFunctionCallNode(stmt) {
			fCall := stmt.(*FunctionCall)
			fCall.Definition = this.findFunctionDecl(prog, fCall.Name)
			if nil == fCall.Definition {
				log.Fatal(fCall.Name, " function not found")
			}
		}
	}
}

func (this RefResolver) findFunctionDecl(prog *Prog, name string) *FunctionDecl {
	for _, stmt := range prog.stmts {
		if IsFunctionDeclNode(stmt) {
			funcDecl := stmt.(*FunctionDecl)
			if funcDecl.Name == name {
				return funcDecl
			}
		}
	}

	return nil
}
