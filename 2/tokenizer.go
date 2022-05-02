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
	return this.Source[this.Pos : this.Pos+1]
}

func (this *CharStream) next() string {
	ch := this.Source[this.Pos : this.Pos+1]
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
	return (this.Pos + 1) >= len(this.Source)
}

/**
 * 词法分析器。
 * 词法分析器的接口像是一个流，词法解析是按需进行的。
 * 支持下面两个操作：
 * next(): 返回当前的Token，并移向下一个Token。
 * peek(): 返回当前的Token，但不移动当前位置。
 */
type Tokenizer struct {
	//Tokens []*Token
	Pos       int //todo remove
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
	//if t.Stream.isEOF() {
	//	return NewToken(EOF, "")
	//} else if {
	//
	//}
	return nil
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
	if str == " " || str == `\n` || str == `\t` {
		return true
	}
	return false
}

//todo remove
func (t *Tokenizer) traceBack(pos int) {}
