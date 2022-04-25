package main

import "fmt"

type Token struct {
	Kind TokenKind
	Text string
}

var tokenArray = []*Token{}

type CharStream struct {
}

func NewCharStream() *CharStream {
	return &CharStream{}
}

func (t *Token) String() string {
	return fmt.Sprintf("%+v", *t)
}

type Tokenizer struct {
	Tokens []*Token
	Pos    int
}

func NewTokenizer(tokens []*Token) *Tokenizer {
	return &Tokenizer{Tokens: tokens}
}

func (t *Tokenizer) Next() *Token {
	if t.Pos >= len(t.Tokens) {
		return nil
	}

	token := t.Tokens[t.Pos]
	t.Pos += 1
	return token
}

func (t *Tokenizer) traceBack(pos int) {
	t.Pos = pos
}
