package main

import (
	"log"
	"testing"
)

var strSignalComment = `//hello,world
china`
var strMultiComments = `/*hello,world
my language
*/
china`

func TestTokenizer_skipMultiComments(t *testing.T) {
	stream := NewCharStream(strMultiComments)
	tokener := NewLexer(stream)
	tokener.skipMultiComments()
	if tokener.Stream.peek() != "\n" {
		t.Error("should be c, but got", tokener.Stream.peek())
	}
}

func TestTokenizer_skipSignalComment(t *testing.T) {
	stream := NewCharStream(strSignalComment)
	tokener := NewLexer(stream)
	tokener.skipSignalComment()
	if tokener.Stream.peek() != "c" {
		t.Error("should be c, but got", tokener.Stream.peek())
	}
}

func TestCharStream_peek(t *testing.T) {
	stream := NewCharStream(strMultiComments)
	c := stream.next()
	if c != "/" {
		t.Error("is not /")
	}
	stream.peek()

	if stream.Pos != 1 {
		t.Error("should be 1")
	}
}

func TestCharStream_next(t *testing.T) {
	stream := NewCharStream(strMultiComments)
	c := stream.next()
	if c != "/" {
		t.Error("is not /")
	}
	stream.next()

	if stream.Pos != 2 {
		t.Error("should be 2")
	}
}

//func TestTokenizer_Next(t *testing.T) {
//	lexer := NewLexer(NewCharStream(source))
//	token := lexer.Peek()
//	for token.Kind != EOF {
//		token = lexer.Next()
//	}
//}

type binnode struct {
	left  *binnode
	right *binnode
	value int
}

func newNode(v int) *binnode {
	return &binnode{value: v}
}

func inOrderRecursion(root *binnode) {
	p := root
	stack := NewStack()
	for p != nil || stack.size != 0 {
		for p != nil {
			stack.Push(p)
			p = p.left
		}

		if stack.size != 0 {
			data := stack.Pop().(*binnode)
			log.Println(data.value)
			p = data.right
		}
	}
}

func preOrderRecursion(root *binnode) {
	p := root
	stack := NewStack()
	for p != nil || stack.size != 0 {
		for p != nil {
			log.Println(p.value)
			stack.Push(p)
			p = p.left
		}

		if stack.size != 0 {
			data := stack.Pop().(*binnode)
			p = data.right
		}
	}
}

func TestCharStream_next33(t *testing.T) {
	root := newNode(2)
	l1 := newNode(1)
	r1 := newNode(4)
	root.left = l1
	root.right = r1

	l1_l := newNode(7)
	l1_r := newNode(8)

	l1.left = l1_l
	l1.right = l1_r

	r1_l := newNode(5)
	r1_r := newNode(6)
	r1.left = r1_l
	r1.right = r1_r

	preOrderRecursion(root)
}
