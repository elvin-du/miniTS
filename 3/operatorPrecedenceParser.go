package main

type OperatorPrecedenceParser struct {
	tokenizer *Tokenizer
	OpStack   *Stack
	AstStack  *Stack
}

func NewOperatorPrecedenceParser(tokenizer *Tokenizer) *OperatorPrecedenceParser {
	return &OperatorPrecedenceParser{tokenizer: tokenizer, OpStack: NewStack(), AstStack: NewStack()}
}

func (this *OperatorPrecedenceParser) Parse() ASTNode {
	for {
		t := this.tokenizer.Next()
		if t.Kind == IntegerLiteral {
			node := NewNode()
			node.SetToken(t)
			node.SetKind(NodeKindScalar)
			node.SetLabel(t.Text)

			this.AstStack.Push(node)
		} else {
			if this.OpStack.Peek() == nil {
				node := NewNode()
				node.SetToken(t)
				node.SetKind(NodeKindOperator)
				node.SetLabel(t.Text)
				this.OpStack.Push(node)
			} else if this.OpStack.Peek().(ASTNode).Token().Precedence() < t.Precedence() {
				node := NewNode()
				node.SetToken(t)
				node.SetKind(NodeKindOperator)
				node.SetLabel(t.Text)
				this.OpStack.Push(node)
			} else {
				for {
					parent := this.OpStack.Pop()
					r := this.AstStack.Pop().(ASTNode)
					l := this.AstStack.Pop().(ASTNode)
					parent.(ASTNode).AddChild(l)
					parent.(ASTNode).AddChild(r)
					this.AstStack.Push(parent)
					if this.OpStack.Peek() == nil ||
						this.OpStack.Peek().(ASTNode).Token().Precedence() < t.Precedence() {
						node := NewNode()
						node.SetToken(t)
						node.SetKind(NodeKindOperator)
						node.SetLabel(t.Text)
						this.OpStack.Push(node)
						break
					}
				}
			}
		}

		if t.Kind == Seperator && t.Text == ";" {
			return this.AstStack.Peek().(ASTNode)
		}
	}
}
