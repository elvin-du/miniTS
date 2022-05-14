package main

import "log"

type RefResolver struct {
	//VisitorAst
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
			this.VisitFunctionDecl(stmt.(*FunctionDecl))
		}
	}
}
func (this *RefResolver) VisitFunctionDecl(funcDecl *FunctionDecl) {
	this.VisitFunctionBody(funcDecl.Body)
}

func (this *RefResolver) VisitFunctionBody(body *FunctionBody) {
	for _, stmt := range body.Stmts {
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

type SymKind int

const (
	Variable  SymKind = 1
	Function          = 2
	Struct            = 3
	Interface         = 4
)

type Symbol struct {
	Name string
	Decl Declaration
	Kind SymKind
}

func NewSymbol(name string, decl Declaration, kind SymKind) *Symbol {
	return &Symbol{Name: name, Decl: decl, Kind: kind}
}

type SymbolTable struct {
}
