package boolean

import (
	"fmt"
	"github.com/avalchev94/go_collections/stack"
	"github.com/avalchev94/go_collections/tree"
	"strings"
)

// Searcher is a struct for boolean evaluation(search).
type Searcher struct {
	tree *tree.Tree
}

// New creates Searcher. Error is returned if the expression has syntax mistakes.
func New(expression string) (*Searcher, error) {
	subTrees := stack.New()
	operators := stack.New()

	reader := newReader(expression)
	for reader.len() > 0 {
		reader.clear(' ')

		if op, err := reader.readOperator(); err == nil {
			switch {
			case op.equal(and) || op.equal(or):
				if operators.Len() > 0 {
					prevOp := operators.Top().(operator)
					if !prevOp.equal(leftBracket) && !op.greater(prevOp) {
						if err := operators.Pop().(operator).createTree(subTrees); err != nil {
							return nil, err
						}
					}
				}
				operators.Push(op)

			case op.equal(leftBracket) || op.equal(not):
				operators.Push(op)

			case op.equal(rightBracket):
				for !operators.Empty() && !operators.Top().(operator).equal(leftBracket) {
					if err := operators.Pop().(operator).createTree(subTrees); err != nil {
						return nil, err
					}
				}
				operators.Pop()
			}
		} else if keyword, err := reader.readKeyword(); err == nil {
			subTrees.Push(tree.New(keyword))
		}
	}

	for operators.Len() > 0 {
		if err := operators.Pop().(operator).createTree(subTrees); err != nil {
			return nil, err
		}
	}

	if subTrees.Len() != 1 {
		return nil, fmt.Errorf("missing operator?")
	}

	return &Searcher{subTrees.Pop().(*tree.Tree)}, nil
}

// Search looks for the keywords in the "searched" string. In the same time,
// it evaluates the boolean expression.
func (s *Searcher) Search(searched string) bool {
	return searchEvaluator(s.tree, searched)
}

func searchEvaluator(t *tree.Tree, searched string) bool {
	switch t.Value.(type) {
	case operator:
		return t.Value.(operator).calculate(
			searchEvaluator(t.Left, searched),
			searchEvaluator(t.Right, searched),
		)

	case string:
		return strings.Contains(searched, t.Value.(string))

	default:
		return false
	}
}
