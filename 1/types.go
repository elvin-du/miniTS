package main

import (
	"fmt"
	"log"
)

type TokenKind int

func (tk TokenKind) String() string {
	switch tk {
	case Keyword:
		return fmt.Sprintf("Keyword")
	case Identifier:
		return fmt.Sprintf("Identifier")
	case StringLiteral:
		return fmt.Sprintf("StringLiteral")
	case Seperator:
		return fmt.Sprintf("Seperator")
	case Operator:
		return fmt.Sprintf("Operator")
	case EOF:
		return fmt.Sprintf("EOF")
	}

	log.Fatalln("tokenkind,", tk, " not define")
	return ""
}

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
