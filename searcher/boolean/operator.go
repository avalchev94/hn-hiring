package boolean

import (
	"fmt"
	"github.com/avalchev94/go_collections/stack"
	"github.com/avalchev94/go_collections/tree"
)

type operator struct {
	char     rune
	priority int8
	arity    int
}

var (
	or           = operator{'|', 0, 2}
	and          = operator{'&', 1, 2}
	not          = operator{'!', 2, 1}
	leftBracket  = operator{'(', 3, 0}
	rightBracket = operator{')', 3, 0}
)

func (op operator) greater(other operator) bool {
	return op.priority > other.priority
}

func (op operator) equal(other operator) bool {
	return op.char == other.char
}

func (op operator) calculate(args ...bool) bool {
	switch {
	case op.equal(not):
		return !args[0]
	case op.equal(and):
		return args[0] && args[1]
	case op.equal(or):
		return args[0] || args[1]

	default:
		return false
	}
}

func (op operator) createTree(subTrees *stack.Stack) error {
	if op.arity <= 0 {
		return fmt.Errorf("operator %c's arity is zero or less", op.char)
	}

	if subTrees.Len() < op.arity {
		return fmt.Errorf("insufficient arguments for operator %c", op.char)
	}

	switch {
	case op.equal(not):
		t := tree.New(op)
		t.Left = subTrees.Pop().(*tree.Tree)
		subTrees.Push(t)

	case op.equal(and) || op.equal(or):
		t := tree.New(op)
		t.Left = subTrees.Pop().(*tree.Tree)
		t.Right = subTrees.Pop().(*tree.Tree)
		subTrees.Push(t)
	}

	return nil
}
