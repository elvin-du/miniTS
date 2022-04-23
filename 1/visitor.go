package main

type VisitorAst struct {
}

func NewVisitorAst() *VisitorAst {
	return &VisitorAst{}
}

func (this VisitorAst) VisitProg(prog *Prog) {
	for _, stmt := range prog.stmts {
		if IsFunctionCallNode(stmt) {
			this.visitFunctionDecl(stmt.(*FunctionCall).Definition)
		} else { //functionDecl
			//do nothing
			this.visitFunctionCall(stmt.(*FunctionCall))
		}
	}
}

func (this VisitorAst) visitFunctionCall(funcDecl *FunctionCall) {
	//do nothing
}

func (this VisitorAst) visitFunctionDecl(funcDecl *FunctionDecl) {
	this.visitFunctionBody(funcDecl.Body)
}

func (this VisitorAst) visitFunctionBody(body *FunctionBody) {
	for _, stmt := range body.Stmts {
		this.visitFunctionCall(stmt.(*FunctionCall))
	}
}
