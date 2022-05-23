package main

import "log"

type Parser struct {
	lexer *Lexer
}

/*
* 当前语法规则：
* prog = statementList? EOF;
* statementList = (variableDecl | functionDecl | expressionStatement)+ ;
* variableDecl : 'let' Identifier typeAnnotation？ ('=' singleExpression) ';';
* typeAnnotation : ':' typeName;
* functionDecl: "function" Identifier "(" ")"  functionBody;
* functionBody : '{' statementList? '}' ;
* statement: functionDecl | expressionStatement;
* expressionStatement: expression ';' ;
* expression: primary (binOP primary)* ;
* primary: StringLiteral | DecimalLiteral | IntegerLiteral | functionCall | '(' expression ')' ;
* binOP: '+' | '-' | '*' | '/' | '=' | '+=' | '-=' | '*=' | '/=' | '==' | '!=' | '<=' | '>=' | '<'
*      | '>' | '&&'| '||'|...;
* functionCall : Identifier '(' parameterList? ')' ;
* parameterList : expression (',' expression)* ;
 */

func NewParser(lexer *Lexer) *Parser {
	return &Parser{lexer: lexer}
}

/**
* 解析Prog
* 语法规则：
* prog = statementList? EOF ;
* statementList =(variableDecl | functionDecl | expressionStatement)+;
 */
func (p *Parser) ParseProg() *Prog {
	return NewProg(p.ParseStatementList())
}

func (p *Parser) ParseStatementList() []Statement {
	stmts := []Statement{}
	t := p.lexer.Peek()
	for t.Kind != EOF {
		stmt := p.ParseStatement()
		if nil != stmt {
			stmts = append(stmts, stmt)
		}
		t = p.lexer.Peek()
	}

	return stmts
}

/**
 * 解析语句。
 * 知识点：在这里，遇到了函数调用和变量赋值，都可能是以Identifier开头的情况，所以预读一个Token是不够的，
 * 所以这里预读了两个Token。
 */
func (this *Parser) ParseStatement() Statement {
	t := this.lexer.Peek()
	if t.Kind == Keyword && t.Text == "function" {
		return this.ParseFunctionDecl()
	} else if t.Text == "let" {
		return this.ParseVariableDecl()
	} else if t.Kind == Identifier {
		//t.Kind == DecimalLiteral ||
		//t.Kind == IntegerLiteral ||
		//t.Kind == StringLiteral ||
		//t.Text == "(" { //todo 没有处理函数调用，变量赋值
		return this.ParseExpressionStmt()
	} else {
		log.Fatalln("cannot recognize a expression:", this.lexer.Peek().Text)
	}

	return nil
}

/**
* 解析函数声明
* 语法规则：
* functionDecl: "function" Identifier "(" parameterList? ")"  functionBody;
* parameterList : Keyword (',' Keyword)* ;
 */
func (p *Parser) ParseFunctionDecl() Statement {
	p.lexer.Next() //跳过function
	param := []string{}

	t := p.lexer.Next()
	if t.Kind == Identifier {
		t1 := p.lexer.Next()
		if t1.Text == "(" {
			t2 := p.lexer.Next()
			for t2.Text != ")" {
				if t2.Kind == Keyword {
					param = append(param, t2.Text)
				}

				t2 = p.lexer.Next()
				if t2.Text == "," {
					t2 = p.lexer.Next()
				}
			}

			if t2.Text == ")" {
				funcBody := p.ParseFunctionBody()
				if nil != funcBody && IsFunctionBodyNode(funcBody) {
					return NewFunctionDecl(t.Text, param, funcBody)
				}
			}
		}
	} else {
		log.Fatal("expect function identifier, but got ", t)
	}

	return nil
}

/**
* 解析函数体
* 语法规则：
* functionBody : '{' functionCall* '}' ;
 */
func (p *Parser) ParseFunctionBody() *Block {
	t := p.lexer.Next()
	stmts := []Statement{}

	if "{" == t.Text {
		for p.lexer.Peek().Kind == Identifier {
			funcCall := p.ParseFunctionCall()
			if nil != funcCall && IsFunctionCallNode(funcCall) {
				stmts = append(stmts, funcCall)
			}
		}

		t = p.lexer.Next()
		if t.Text == "}" {
			return NewBlock(stmts)
		} else {
			log.Fatal("expect },but got ", t.Text)
		}
	} else {
		log.Fatal("expect {, but got ", t.Text)
	}

	return nil
}

/**
* 解析函数调用
* 语法规则：
* functionCall : Identifier '(' parameterList? ')' ;
* parameterList : StringLiteral (',' StringLiteral)* ;
 */
func (p *Parser) ParseFunctionCall() Statement {
	var parameters []string
	t := p.lexer.Next()

	if t.Kind == Identifier {
		t1 := p.lexer.Next()
		if t1.Text == "(" {
			t2 := p.lexer.Next()
			for t2.Text != ")" {
				if t2.Kind == StringLiteral {
					parameters = append(parameters, t2.Text)
				} else {
					log.Println("Expecting parameter in FunctionCall, while we got a ", t2.Text)
					return nil
				}
				t2 = p.lexer.Next()
				if t2.Text != ")" {
					if t2.Text == "," {
						t2 = p.lexer.Next() //消化掉,
					} else {
						log.Println("Expecting a comma , in FunctionCall, while we got a " + t2.Text)
						return nil
					}
				}
			}
			t2 = p.lexer.Next() //消化掉 ;
			if t2.Text == ";" {
				return NewFunctionCall(t.Text, parameters)
			} else {
				log.Println("Expecting a comma ; in FunctionCall, while we got a " + t2.Text)
				return nil
			}
		}
	}

	return nil
}

/**
 * 解析变量声明
 * 语法规则：
 * variableDecl : 'let' Identifier typeAnnotation？ ('=' singleExpression)? ';';
* typeAnnotation : ':' typeName;
*/
func (this *Parser) ParseVariableDecl() Statement {
	this.lexer.Next() //skip let token
	t := this.lexer.Next()
	if t.Kind == Identifier {
		name := t.Text
		typ := ""
		var init Expression = nil

		t = this.lexer.Next()
		if t.Text == ":" {
			t = this.lexer.Next()
			if t.Kind == Identifier {
				typ = t.Text
			} else {
				log.Fatalln("kind should be indentifier")
			}
		}

		//初始化部分
		if t.Text == "=" {
			init = this.ParseExpression()
		}

		t = this.lexer.Peek()
		if t.Text == ";" {
			return NewVariableDecl(name, typ, init)
		} else {
			log.Fatalln("kind should be ; ")
		}
	}

	return nil
}

func (this *Parser) ParseExpression() Expression {
	return this.ParseBinaryExpr()
}

// b =  a + 1 * 3;,这里只处理后面的表达式，也就是 a+1*3的部分
func (this *Parser) ParseBinaryExpr() Expression {
	OpStack := NewStack()
	AstStack := NewStack()

	for {
		t := this.lexer.Next()
		//todo 现在只能处理int类型 和变量
		if t.Kind == IntegerLiteral || t.Kind == Identifier {
			node := NewNode()
			node.SetToken(t)
			node.SetKind(NodeKindScalar)
			node.SetLabel(t.Text)

			AstStack.Push(node)
		} else if t.Text == "(" { //todo 现在还无法处理 a*(2+3)的情况，
			expr := this.ParseExpression()
			AstStack.Push(expr)
			this.lexer.Next() //skip ")"符号
		} else {
			//如果操作符栈为空的话，直接生成AST节点，并压入操作符栈
			if OpStack.Peek() == nil {
				node := NewNode()
				node.SetToken(t)
				node.SetKind(NodeKindOperator)
				node.SetLabel(t.Text)
				OpStack.Push(node)
			} else if OpStack.Peek().(ASTNode).Token().Precedence() < t.Precedence() {
				node := NewNode()
				node.SetToken(t)
				node.SetKind(NodeKindOperator)
				node.SetLabel(t.Text)
				OpStack.Push(node)
			} else {
				for {
					parent := OpStack.Pop()
					r := AstStack.Pop().(ASTNode)
					l := AstStack.Pop().(ASTNode)
					parent.(ASTNode).AddChild(l)
					parent.(ASTNode).AddChild(r)
					AstStack.Push(parent)
					if OpStack.Peek() == nil ||
						OpStack.Peek().(ASTNode).Token().Precedence() < t.Precedence() {
						node := NewNode()
						node.SetToken(t)
						node.SetKind(NodeKindOperator)
						node.SetLabel(t.Text)
						OpStack.Push(node)
						break
					}
				}
			}
		}

		if t.Kind == Seperator && t.Text == ";" {
			return AstStack.Peek().(ASTNode)
		}
	}

	log.Fatalln("should be here")
	return nil
}

func (this *Parser) ParseExpressionStmt() Statement {
	expr := this.ParseExpression()
	if nil != expr {
		t := this.lexer.Peek()
		if t.Text == ";" {
			return NewExpressionStmt(expr)
		} else {
			log.Fatalln("should be ; but got:", t.Text)
		}
	} else {
		log.Fatalln("expect expression, bug got nil")
	}

	return nil
}
