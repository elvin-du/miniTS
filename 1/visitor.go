package main

//
//type VisitorAst struct {
//}
//
//func NewVisitorAst() *VisitorAst {
//	return &VisitorAst{}
//}
//
//func (this *VisitorAst) VisitProg(prog *Prog) {
//	for _, stmt := range prog.stmts {
//		if IsFunctionCallNode(stmt) {
//			this.VisitFunctionDecl(stmt.(*FunctionCall).Definition)
//		} else { //functionDecl
//			this.VisitFunctionCall(stmt.(*FunctionCall))
//		}
//	}
//}
//
//func (this *VisitorAst) VisitFunctionCall(funcDecl *FunctionCall) {
//	//do nothing
//}
//
//func (this *VisitorAst) VisitFunctionDecl(funcDecl *FunctionDecl) {
//	this.VisitFunctionBody(funcDecl.Body)
//}
//
//func (this *VisitorAst) VisitFunctionBody(body *FunctionBody) {
//	for _, stmt := range body.Stmts {
//		this.VisitFunctionCall(stmt.(*FunctionCall))
//	}
//}
