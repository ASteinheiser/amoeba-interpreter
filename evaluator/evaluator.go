package evaluator

import (
	"github.com/ASteinheiser/amoeba-interpreter/ast"
	"github.com/ASteinheiser/amoeba-interpreter/object"
)

// Eval will evaluate a program
func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	// Statements
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	// Expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}

	return nil
}

func evalStatements(stmts []ast.Statement) (result object.Object) {
	for _, statement := range stmts {
		result = Eval(statement)
	}
	return
}
