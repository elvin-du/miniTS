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
	case IntegerLiteral:
		return fmt.Sprintf("IntegerLiteral")
	case Seperator:
		return fmt.Sprintf("Seperator")
	case Operator:
		return fmt.Sprintf("Operator")
	case EOF:
		return fmt.Sprintf("EOF")
	}

	log.Fatalf("tokenkind %d,%s", tk, " not define")
	return ""
}

const (
	Keyword        TokenKind = 1
	Identifier               = 2
	StringLiteral            = 3
	IntegerLiteral           = 7
	DecimalLiteral           = 8
	Seperator                = 4
	Operator                 = 5
	EOF                      = 6
)

type NodeKind int

const (
	NodeKindProg         NodeKind = 1
	NodeKineBinaryExpr            = 6
	NodeKindStatement             = 2
	NodeKindScalar                = 7
	NodeKindOperator              = 9
	NodeKindVariable              = 8
	NodeKindFunctionDecl          = 3
	NodeKindFunctionCall          = 4
	NodeKindFunctionBody          = 5
)
