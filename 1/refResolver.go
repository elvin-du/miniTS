package main

import "log"

type RefResolver struct {
	VisitorAst
	prog *Prog
}

func NewRefResolver() *RefResolver {
	return &RefResolver{}
}

func (this *RefResolver) VisitProg(prog *Prog) {
	this.prog = prog
	for _, stmt := range prog.stmts {
		if IsFunctionCallNode(stmt) {
			this.resolveFunctionCall(prog, stmt.(*FunctionCall))
		} else { //functionDecl
			this.visitFunctionDecl(stmt.(*FunctionDecl))
		}
	}
}

func (this *RefResolver) VisitFunctionBody(fBody *FunctionBody) {
	for _, stmt := range fBody.Stmts {
		this.resolveFunctionCall(this.prog, stmt.(*FunctionCall))
	}
}

func (this *RefResolver) resolveFunctionCall(prog *Prog, fCall *FunctionCall) {
	fDecl := this.findFunctionDecl(prog, fCall.Name)
	if nil != fDecl {
		fCall.Definition = fDecl
		return
	} else {
		if fCall.Name != "println" {
			log.Fatalln("function", fCall.Name, "not found")
		}
	}
}

func (this *RefResolver) findFunctionDecl(prog *Prog, name string) *FunctionDecl {
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
