package ast

import (
	"interpreter/tokens"
	"strconv"
)

type Node interface {
	String() string
}

type Binary struct {
	Token tokens.Token
	Left  Node
	Right Node
}

func (b *Binary) String() string {
	return b.Left.String() + tokens.TYPE_MAP[b.Token.Type] + b.Right.String() + "\n"
}

type Unary struct {
	Token tokens.Token
	Right Node
}

func (u *Unary) String() string {
	return tokens.TYPE_MAP[u.Token.Type] + u.Right.String() + "\n"
}

type Number struct {
	Token tokens.Token
	Value float64
}

func (n *Number) String() string {
	return strconv.FormatFloat(n.Value, 'g', -1, 64)
}
