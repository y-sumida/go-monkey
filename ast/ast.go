package ast

import "go-monkey/token"

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (ls *LetStatement) statomentNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
