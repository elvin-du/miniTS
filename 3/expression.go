package main

type Expr struct {
	*node
}

func NewExpr(kind NodeKind, token *Token) *Expr {
	expr := &Expr{NewNode()}
	expr.SetLabel(token.Text)
	expr.SetKind(kind)
	expr.SetToken(token)
	return expr
}
