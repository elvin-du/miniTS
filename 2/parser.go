package main

import "log"

type Parser struct {
	tokenizer *Tokenizer
}

func NewParser(tokenizer *Tokenizer) *Parser {
	return &Parser{tokenizer: tokenizer}
}

/**
 * 解析Prog
 * 语法规则：
 * prog = (functionDecl | functionCall)* ;
 */
func (p *Parser) ParseProg() *Prog {
	stmts := []Statement{}
	for {
		stmt := p.ParseFunctionDecl()
		if nil != stmt {
			stmts = append(stmts, stmt)
			continue
		}

		stmt = p.ParseFunctionCall()
		if nil != stmt {
			stmts = append(stmts, stmt)
			continue
		}

		if nil == stmt {
			break
		}
	}

	return NewProg(stmts)
}

/**
 * 解析函数声明
 * 语法规则：
 * functionDecl: "function" Identifier "(" parameterList? ")"  functionBody;
 * parameterList : Keyword (',' Keyword)* ;
 */
func (p *Parser) ParseFunctionDecl() Statement {
	t := p.tokenizer.Next()
	param := []string{}

	if t.Kind == Keyword && t.Text == "function" {
		t = p.tokenizer.Next()
		if t.Kind == Identifier {
			t1 := p.tokenizer.Next()
			if t1.Text == "(" {
				t2 := p.tokenizer.Next()
				for t2.Text != ")" {
					if t2.Kind == Keyword {
						param = append(param, t2.Text)
					}

					t2 = p.tokenizer.Next()
					if t2.Text == "," {
						t2 = p.tokenizer.Next()
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
	}

	return nil
}

/**
 * 解析函数体
 * 语法规则：
 * functionBody : '{' functionCall* '}' ;
 */
func (p *Parser) ParseFunctionBody() *FunctionBody {
	t := p.tokenizer.Next()
	stmts := []Statement{}

	if "{" == t.Text {
		for {
			stmt := p.ParseFunctionCall()
			if nil != stmt && IsFunctionCallNode(stmt) {
				stmts = append(stmts, stmt)
			} else {
				break
			}
		}

		t = p.tokenizer.Next()
		if t.Text == "}" {
			return NewFunctionBody(stmts)
		} else {
			log.Fatal("expect },but got ", t.Text)
		}
	} else {
		log.Println("expect {, but got ", t.Text)
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
	parameters := []string{}
	t := p.tokenizer.Next()

	if t.Kind == Identifier {
		t1 := p.tokenizer.Next()
		if t1.Text == "(" {
			t2 := p.tokenizer.Next()
			for t2.Text != ")" {
				if t2.Kind == StringLiteral {
					parameters = append(parameters, t2.Text)
				} else {
					log.Println("Expecting parameter in FunctionCall, while we got a ", t2.Text)
					return nil
				}
				t2 = p.tokenizer.Next()
				if t2.Text != ")" {
					if t2.Text == "," {
						t2 = p.tokenizer.Next()
					} else {
						log.Println("Expecting a comma in FunctionCall, while we got a " + t2.Text)
						return nil
					}
				}
			}
			t2 = p.tokenizer.Next()
			if t2.Text == ";" {
				return NewFunctionCall(t.Text, parameters)
			} else {
				log.Println("Expecting a comma in FunctionCall, while we got a " + t2.Text)
				return nil
			}
		}
	}

	return nil
}
