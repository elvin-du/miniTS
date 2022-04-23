package main

func IsStatementNode(node AstNode) bool {
	k := node.Kind()
	if k == NodeKindFunctionDecl || k == NodeKindFunctionCall {
		return true
	}

	return false
}

func IsFunctionCallNode(node AstNode) bool {
	k := node.Kind()
	if k == NodeKindFunctionCall {
		return true
	}
	return false
}

func IsFunctionDeclNode(node AstNode) bool {
	k := node.Kind()
	if k == NodeKindFunctionDecl {
		return true
	}
	return false
}

func IsFunctionBodyNode(node AstNode) bool {
	k := node.Kind()
	if k == NodeKindFunctionBody {
		return true
	}
	return false
}
