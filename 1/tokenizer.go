package main

type Token struct {
	Kind TokenKind
	Text string
}

// 一个Token数组，代表了下面这段程序做完词法分析后的结果：
/*
//一个函数的声明，这个函数很简单，只打印"Hello World!"
function sayHello(){
    println("Hello World!");
}
//调用刚才声明的函数
sayHello();
*/
var tokenArray = []Token{
	{Kind: Keyword, Text: "function"},
	{Kind: Identifier, Text: "sayHello"},
	{Kind: Seperator, Text: "("},
	{Kind: Seperator, Text: ")"},
	{Kind: Seperator, Text: "{"},
	{Kind: Identifier, Text: "println"},
	{Kind: Seperator, Text: "("},
	{Kind: StringLiteral, Text: "Hello World!"},
	{Kind: Seperator, Text: ")"},
	{Kind: Seperator, Text: ";"},
	{Kind: Seperator, Text: "}"},
	{Kind: Identifier, Text: "sayHello"},
	{Kind: Seperator, Text: "("},
	{Kind: Seperator, Text: ")"},
	{Kind: Seperator, Text: ";"},
	{Kind: EOF, Text: ""},
}

type Tokenizer struct {
	Tokens []*Token
	Pos    int
}

func NewTokenizer() *Tokenizer {
	return &Tokenizer{}
}

func (t *Tokenizer) Next() *Token {
	if t.Pos >= len(t.Tokens) {
		return nil
	}

	token := t.Tokens[t.Pos]
	t.Pos += t.Pos
	return token
}

func (t *Tokenizer) traceBack(pos int) {
	t.Pos = pos
}
