package main

import "testing"

var strSignalComment = `//hello,world
china`
var strMultiComments = `/*hello,world
my language
*/
china`

func TestTokenizer_skipMultiComments(t *testing.T) {
	stream := NewCharStream(strMultiComments)
	tokener := NewTokenizer(stream)
	tokener.skipMultiComments()
	if tokener.Stream.peek() != "\n" {
		t.Error("should be c, but got", tokener.Stream.peek())
	}
}

func TestTokenizer_skipSignalComment(t *testing.T) {
	stream := NewCharStream(strSignalComment)
	tokener := NewTokenizer(stream)
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

func TestTokenizer_Next(t *testing.T) {
	tokenizer := NewTokenizer(NewCharStream(source))
	token := tokenizer.Peek()
	for token.Kind != EOF {
		token = tokenizer.Next()
	}
}
