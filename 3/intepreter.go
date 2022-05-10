package main

//
//import "log"
//
//type Intepreter struct {
//}
//
//func NewIntepreter() *Intepreter {
//	return &Intepreter{}
//}
//
//func (this *Intepreter) Start(prog *Prog) {
//	for _, stmt := range prog.stmts {
//		if IsFunctionCallNode(stmt) {
//			this.RunFunction(stmt.(*FunctionCall))
//		}
//	}
//}
//
//func (this *Intepreter) RunFunction(fCall *FunctionCall) {
//	if fCall.Name == "println" {
//		log.Println(fCall.Parameters)
//		return
//	}
//
//	if fCall.Definition != nil {
//		this.VisitFunctionBody(fCall.Definition.Body)
//	}
//}
//
//func (this *Intepreter) VisitFunctionBody(fBody *FunctionBody) {
//	for _, stmt := range fBody.Stmts {
//		this.RunFunction(stmt.(*FunctionCall))
//	}
//}
