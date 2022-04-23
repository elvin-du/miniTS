package main

type TokenKind int

const (
	Keyword       TokenKind = 1
	Identifier              = 2
	StringLiteral           = 3
	Seperator               = 4
	Operator                = 5
	EOF                     = 6
)

type NodeKind int

const (
	NodeKindProg         NodeKind = 1
	NodeKindStatement             = 2
	NodeKindFunctionDecl          = 3
	NodeKindFunctionCall          = 4
	NodeKindFunctionBody          = 5
)
