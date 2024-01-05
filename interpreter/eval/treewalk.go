package eval

import (
	"interpreter/ast"
	"interpreter/tokens"
)

func TreeWalk(node ast.Node) float64 {
	switch node := node.(type) {
	case *ast.Binary:
		lhs := TreeWalk(node.Left)
		rhs := TreeWalk(node.Right)
		switch node.Token.Type {
		case tokens.PLUS:
			return lhs + rhs
		case tokens.MINUS:
			return lhs - rhs
		case tokens.ASTERIKS:
			return lhs * rhs
		case tokens.SLASH:
			return lhs / rhs
		}
	case *ast.Unary:
		return TreeWalk(node.Right) * -1
	case *ast.Number:
		return node.Value
	}
	return 0
}
