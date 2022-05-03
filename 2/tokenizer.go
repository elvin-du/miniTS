package main

import (
	"fmt"
	"log"
)

type Token struct {
	Kind TokenKind
	Text string
}

func NewToken(kind TokenKind, text string) *Token {
	return &Token{Kind: kind, Text: text}
}

//var tokenArray = []*Token{}

func (t *Token) String() string {
	return fmt.Sprintf("%+v", *t)
}

/*
 * 一个字符串流。其操作为：
 * peek():预读下一个字符，但不移动指针；
 * next():读取下一个字符，并且移动指针；
 * isEOF():判断是否已经到了结尾。s
 */
type CharStream struct {
	Source string
	Pos    int
	Line   int //只是用于日志，debug等
	Col    int //只是用于日志，debug等
}

func NewCharStream(source string) *CharStream {
	return &CharStream{Source: source, Pos: 0, Line: 1, Col: 0}
}

func (this *CharStream) peek() string {
	if this.isEOF() {
		return ""
	}

	return string(this.Source[this.Pos])
}

func (this *CharStream) next() string {
	if this.isEOF() {
		return ""
	}

	ch := string(this.Source[this.Pos])
	if ch == "\n" {
		this.Line++
		this.Col = 0
	} else {
		this.Col++
	}
	this.Pos++

	return ch
}

func (this *CharStream) isEOF() bool {
	return this.Pos >= len(this.Source)
}

/**
 * 词法分析器。
 * 词法分析器的接口像是一个流，词法解析是按需进行的。
 * 支持下面两个操作：
 * next(): 返回当前的Token，并移向下一个Token。
 * peek(): 返回当前的Token，但不移动当前位置。
 */
type Tokenizer struct {
	Stream    *CharStream
	NextToken *Token
}

func NewTokenizer(stream *CharStream) *Tokenizer {
	return &Tokenizer{Stream: stream, NextToken: &Token{Kind: EOF, Text: ""}}
}

func (t *Tokenizer) Next() *Token {
	if t.NextToken.Kind == EOF && !t.Stream.isEOF() {
		t.NextToken = t.getAToken()
	}
	lastToken := t.NextToken

	t.NextToken = t.getAToken()
	log.Printf("token:%+v", lastToken)
	return lastToken
}

func (t *Tokenizer) Peek() *Token {
	if t.NextToken.Kind == EOF && !t.Stream.isEOF() {
		t.NextToken = t.getAToken()
	}
	return t.NextToken
}

func (t *Tokenizer) getAToken() *Token {
	t.skipSpaces()
	if t.Stream.isEOF() {
		return NewToken(EOF, "")
	}

	c := t.Stream.peek()
	switch c {
	case "(", ")", "{", "}", ";", ",":
		t.Stream.next()
		return NewToken(Seperator, c)
	case `"`:
		return t.parseStringLiteral()
	case "/":
		t.Stream.next() //跳过第一个/
		c1 := t.Stream.peek()
		if "*" == c1 {
			t.skipMultiComments()
			return t.getAToken()
		} else if "/" == c1 {
			t.skipSignalComment()
			return t.getAToken()
		} else if "=" == c1 {
			t.Stream.next()
			return NewToken(Operator, "/=")
		} else {
			return NewToken(Operator, "/")
		}
	case "*":
		t.Stream.next() //跳过*
		c1 := t.Stream.peek()
		if "=" == c1 {
			t.Stream.next()
			return NewToken(Operator, "*=")
		} else {
			return NewToken(Operator, "*")
		}
	case "-":
		t.Stream.next()
		c1 := t.Stream.next()
		if "=" == c1 {
			t.Stream.next()
			return NewToken(Operator, "-=")
		} else {
			return NewToken(Operator, "-")
		}
	case "+":
		t.Stream.next()
		c1 := t.Stream.peek()
		if "=" == c1 {
			t.Stream.next()
			return NewToken(Operator, "+=")
		} else {
			return NewToken(Operator, "+")
		}
	}

	if t.isLetter(c) {
		return t.parseIdentifier()
	}

	//暂时去掉不能识别的字符
	log.Println("Unrecognized pattern meeting ': ", c, "', at", t.Stream.Line, " col: ", t.Stream.Col)
	t.Stream.next() //skip unrecognized char
	return t.getAToken()
}

/**
 * 字符串字面量。
 * 目前只支持双引号，并且不支持转义。
 */
func (t *Tokenizer) parseStringLiteral() *Token {
	token := NewToken(StringLiteral, "")
	t.Stream.next() //去掉"

	for !t.Stream.isEOF() {
		if `"` == t.Stream.peek() {
			t.Stream.next() //去掉"
			return token
		}

		token.Text += t.Stream.next()
	}

	log.Fatal("should not be here")
	return nil
}

/**
 * 解析标识符。从标识符中还要挑出关键字。
 */
func (t *Tokenizer) parseIdentifier() *Token {
	token := NewToken(Identifier, "")
	//第一个字符不用判断，因为在调用者那里已经判断过了
	token.Text += t.Stream.next()

	for !t.Stream.isEOF() && t.isLetterDigitOrUnderScore(t.Stream.peek()) {
		token.Text += t.Stream.next()
	}

	if token.Text == "function" {
		token.Kind = Keyword
	}

	return token
}
func (t *Tokenizer) skipSignalComment() {
	//跳过第二个/，第一个之前已经跳过去了。
	t.Stream.next()
	for !t.Stream.isEOF() {
		if t.Stream.next() == "\n" {
			return
		}
	}
}
func (t *Tokenizer) skipMultiComments() {
	//跳过*，之前的/已经跳过去了。
	t.Stream.next()

	for !t.Stream.isEOF() {
		ch1 := t.Stream.next()
		if ch1 == "*" {
			if t.Stream.peek() == "/" {
				t.Stream.next() //跳过最后一个/
				return
			}
		}
	}

	log.Fatalf("multicoments in invalid. not found */, line:%d,column:%d", t.Stream.Line, t.Stream.Col)
}

func (t *Tokenizer) skipSpaces() {
	for t.isSpace(t.Stream.peek()) {
		t.Stream.next()
	}
}
func (t *Tokenizer) isSpace(str string) bool {
	if str == " " || str == "\n" || str == "\t" {
		return true
	}
	return false
}

func (t *Tokenizer) isDigit(str string) bool {
	if len(str) > 0 {
		return str[0] >= '0' && str[0] <= '9'
	}
	return false
}

func (t *Tokenizer) isLetter(str string) bool {
	if len(str) > 0 {
		c := str[0]
		return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z'
	}
	return false
}
func (t *Tokenizer) isLetterDigitOrUnderScore(str string) bool {
	if len(str) > 0 {
		c := str[0]
		return c >= 'a' && c <= 'z' ||
			c >= 'A' && c <= 'Z' ||
			c >= '0' && c <= '9' ||
			c == '_'
	}
	return false
}
