package boolean

import (
	"github.com/avalchev94/go_collections/stack"
	"github.com/avalchev94/go_collections/tree"
	"strings"
)

// Searcher is a struct for boolean evaluation(search).
type Searcher struct {
	tree *tree.Tree
}

// Error struct/type is returned when new Searcher is created.
// It has json bindings, because it is supposed to be send to client.
type Error struct {
	Column int    `json:"column"`
	Error  string `json:"error"`
}

// New creates Searcher. Error is returned if the expression has syntax mistakes.
func New(expression string) (*Searcher, *Error) {
	subTrees := stack.New()
	operators := stack.New()

	reader := newReader(expression)
	for reader.len() > 0 {
		reader.clear(' ')

		token, err := reader.readToken()
		if token == nil {
			return nil, createError(reader, err.Error())
		}

		switch token.(type) {
		// operator
		case operator:
			op := token.(operator)
			switch {
			case op.equal(and) || op.equal(or):
				if operators.Len() > 0 {
					prevOp := operators.Top().(operator)
					if !prevOp.equal(leftBracket) && !op.greater(prevOp) {
						if err := operators.Pop().(operator).createTree(subTrees); err != nil {
							return nil, createError(reader, err.Error())
						}
					}
				}
				operators.Push(op)

			case op.equal(leftBracket) || op.equal(not):
				operators.Push(op)

			case op.equal(rightBracket):
				for {
					if operators.Empty() {
						return nil, createError(reader, "brackets does not match")
					}

					if op := operators.Pop().(operator); op.equal(leftBracket) {
						break
					} else {
						if err := op.createTree(subTrees); err != nil {
							return nil, createError(reader, err.Error())
						}
					}
				}
			}

		// keyword
		case string:
			subTrees.Push(tree.New(token.(string)))
		}
	}

	for operators.Len() > 0 {
		op := operators.Pop().(operator)

		if op.equal(leftBracket) {
			return nil, createError(nil, "brackets does not match")
		} else if err := op.createTree(subTrees); err != nil {
			return nil, createError(nil, err.Error())
		}
	}

	if subTrees.Len() != 1 {
		return nil, createError(nil, "missing operator?")
	}

	return &Searcher{subTrees.Pop().(*tree.Tree)}, nil
}

func createError(r *reader, err string) *Error {
	if r != nil {
		return &Error{r.currentIndex(), err}
	}

	if err != "" {
		return &Error{-1, err}
	}

	return &Error{-1, "undefined?"}
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
