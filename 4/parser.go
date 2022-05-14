package main

//
//import "log"
//
//type Parser struct {
//	tokenizer *Tokenizer
//}
//
//func NewParser(tokenizer *Tokenizer) *Parser {
//	return &Parser{tokenizer: tokenizer}
//}
//
///**
// * 解析Prog
// * 语法规则：
// * prog = (functionDecl | functionCall)* ;
// */
//func (p *Parser) ParseProg() *Prog {
//	stmts := []Statement{}
//	t := p.tokenizer.Peek()
//	for t.Kind != EOF {
//		var stmt Statement
//		if t.Kind == Keyword && "function" == t.Text {
//			stmt = p.ParseFunctionDecl()
//
//		} else if t.Kind == Identifier {
//			stmt = p.ParseFunctionCall()
//		}
//
//		if nil != stmt {
//			stmts = append(stmts, stmt)
//		}
//		t = p.tokenizer.Peek()
//	}
//
//	return NewProg(stmts)
//}
//
///**
// * 解析函数声明
// * 语法规则：
// * functionDecl: "function" Identifier "(" parameterList? ")"  functionBody;
// * parameterList : Keyword (',' Keyword)* ;
// */
//func (p *Parser) ParseFunctionDecl() Statement {
//	p.tokenizer.Next() //跳过function
//	param := []string{}
//
//	t := p.tokenizer.Next()
//	if t.Kind == Identifier {
//		t1 := p.tokenizer.Next()
//		if t1.Text == "(" {
//			t2 := p.tokenizer.Next()
//			for t2.Text != ")" {
//				if t2.Kind == Keyword {
//					param = append(param, t2.Text)
//				}
//
//				t2 = p.tokenizer.Next()
//				if t2.Text == "," {
//					t2 = p.tokenizer.Next()
//				}
//			}
//
//			if t2.Text == ")" {
//				funcBody := p.ParseFunctionBody()
//				if nil != funcBody && IsFunctionBodyNode(funcBody) {
//					return NewFunctionDecl(t.Text, param, funcBody)
//				}
//			}
//		}
//	} else {
//		log.Fatal("expect function identifier, but got ", t)
//	}
//
//	return nil
//}
//
///**
// * 解析函数体
// * 语法规则：
// * functionBody : '{' functionCall* '}' ;
// */
//func (p *Parser) ParseFunctionBody() *FunctionBody {
//	t := p.tokenizer.Next()
//	stmts := []Statement{}
//
//	if "{" == t.Text {
//		for p.tokenizer.Peek().Kind == Identifier {
//			funcCall := p.ParseFunctionCall()
//			if nil != funcCall && IsFunctionCallNode(funcCall) {
//				stmts = append(stmts, funcCall)
//			}
//		}
//
//		t = p.tokenizer.Next()
//		if t.Text == "}" {
//			return NewFunctionBody(stmts)
//		} else {
//			log.Fatal("expect },but got ", t.Text)
//		}
//	} else {
//		log.Fatal("expect {, but got ", t.Text)
//	}
//
//	return nil
//}
//
///**
// * 解析函数调用
// * 语法规则：
// * functionCall : Identifier '(' parameterList? ')' ;
// * parameterList : StringLiteral (',' StringLiteral)* ;
// */
//func (p *Parser) ParseFunctionCall() Statement {
//	var parameters []string
//	t := p.tokenizer.Next()
//
//	if t.Kind == Identifier {
//		t1 := p.tokenizer.Next()
//		if t1.Text == "(" {
//			t2 := p.tokenizer.Next()
//			for t2.Text != ")" {
//				if t2.Kind == StringLiteral {
//					parameters = append(parameters, t2.Text)
//				} else {
//					log.Println("Expecting parameter in FunctionCall, while we got a ", t2.Text)
//					return nil
//				}
//				t2 = p.tokenizer.Next()
//				if t2.Text != ")" {
//					if t2.Text == "," {
//						t2 = p.tokenizer.Next() //消化掉,
//					} else {
//						log.Println("Expecting a comma , in FunctionCall, while we got a " + t2.Text)
//						return nil
//					}
//				}
//			}
//			t2 = p.tokenizer.Next() //消化掉 ;
//			if t2.Text == ";" {
//				return NewFunctionCall(t.Text, parameters)
//			} else {
//				log.Println("Expecting a comma ; in FunctionCall, while we got a " + t2.Text)
//				return nil
//			}
//		}
//	}
//
//	return nil
//}

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
