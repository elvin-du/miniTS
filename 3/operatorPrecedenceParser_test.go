package main

import (
	"testing"
)

func TestOperatorPrecedenceParser_Parse(t *testing.T) {
	expr := `2+3*5/6-7;`
	tokenizer := NewTokenizer(NewCharStream(expr))
	parser := NewOperatorPrecedenceParser(tokenizer)
	node := parser.Parse()
	t.Log(node)
}
