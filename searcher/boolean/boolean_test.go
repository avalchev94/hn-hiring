package boolean

import (
	"github.com/avalchev94/go_collections/stack"
	"github.com/avalchev94/go_collections/tree"
	"github.com/stretchr/testify/assert"
	"testing"
)

func AssertTokenString(t *testing.T, r *reader, actual string) {
	assert := assert.New(t)

	token, err := r.readToken()
	assert.NotNil(token)
	assert.Nil(err)
	assert.IsType(token, "")
	assert.Equal(token.(string), actual)
}

func AssertTokenOperator(t *testing.T, r *reader, actual operator) {
	assert := assert.New(t)

	token, err := r.readToken()
	assert.NotNil(token)
	assert.Nil(err)
	assert.IsType(token, actual)
	assert.True(token.(operator).equal(actual))
}

func AssertInvalidToken(t *testing.T, r *reader) {
	assert := assert.New(t)

	token, err := r.readToken()
	assert.Nil(token)
	assert.NotNil(err)
}

func TestReader(t *testing.T) {
	assert := assert.New(t)

	r := newReader("  \"A\"|\"Bac12\"&  !(\"1\"#\"A\")\"\" \"ABC")
	assert.NotNil(r)

	assert.NoError(r.clear(' '))
	assert.Equal(r.len(), 31)
	assert.Equal(r.currentIndex(), 2)

	AssertTokenString(t, r, "A")
	AssertTokenOperator(t, r, or)
	AssertTokenString(t, r, "Bac12")
	AssertTokenOperator(t, r, and)

	assert.NoError(r.clear(' '))
	AssertTokenOperator(t, r, not)
	AssertTokenOperator(t, r, leftBracket)
	AssertTokenString(t, r, "1")
	AssertInvalidToken(t, r)
	AssertTokenString(t, r, "A")
	AssertTokenOperator(t, r, rightBracket)

	AssertInvalidToken(t, r)
	assert.NoError(r.clear(' '))
	AssertInvalidToken(t, r)

	// stream ended asserts
	_, err := r.read()
	assert.NotNil(err)
	AssertInvalidToken(t, r)
}

func TestOperator(t *testing.T) {
	assert := assert.New(t)

	assert.True(and.greater(or))
	assert.True(not.greater(and))
	assert.True(and.equal(and))

	assert.True(and.calculate(true, true))
	assert.False(and.calculate(true, false))
	assert.True(or.calculate(true, false))
	assert.False(or.calculate(false, false))
	assert.True(not.calculate(false))
	assert.False(leftBracket.calculate()) /// calculate works only for and, or, not

	s := stack.New()
	s.Push(tree.New("ABC"))
	// handle the errors first
	assert.Error(and.createTree(s))
	assert.Error(or.createTree(s))
	assert.Error(leftBracket.createTree(s))

	assert.NoError(not.createTree(s))
	assert.Equal(s.Len(), 1)
	assert.True(s.Top().(*tree.Tree).Value.(operator).equal(not))
	assert.Equal(s.Top().(*tree.Tree).Left.Value.(string), "ABC")
	assert.Nil(s.Top().(*tree.Tree).Right)

	s.Push(tree.New("CBA"))
	assert.NoError(and.createTree(s))
	assert.Equal(s.Len(), 1)
	assert.True(s.Top().(*tree.Tree).Value.(operator).equal(and))
}

//func TestOperator(t *testing.T) {
//	if or.greater(and) || and.greater(not) || not.greater(leftBracket) || leftBracket.greater(rightBracket) {
//		t.Error("greater's evaluation is incorrect")
//	}
//
//	if or.equal(and) || and.equal(not) || not.equal(leftBracket) || leftBracket.equal(rightBracket) {
//		t.Error("equal's evaluation is incorrect")
//	}
//
//	parameters := stack.New()
//	if or.calculate(parameters) == nil ||
//		and.calculate(parameters) == nil ||
//		not.calculate(parameters) == nil ||
//		leftBracket.calculate(parameters) == nil ||
//		rightBracket.calculate(parameters) == nil {
//		t.Error("parameters stack is empty. calculation should return error")
//	}
//
//	parameters.Push(false)
//	parameters.Push(true)
//	if or.calculate(parameters) != nil || parameters.Len() != 1 || parameters.Top().(bool) != true {
//		t.Error("calculation of operator or failed")
//	}
//
//	if not.calculate(parameters) != nil || parameters.Len() != 1 || parameters.Top().(bool) != false {
//		t.Error("calculation of operator not failed")
//	}
//
//	parameters.Push(true)
//	if and.calculate(parameters) != nil || parameters.Len() != 1 || parameters.Top().(bool) != false {
//		t.Error("calculation of operator and failed")
//	}
//
//	otherOp := operator{'<', 5, 1}
//	if otherOp.calculate(parameters) == nil {
//		t.Error("calculation for user created operators should fail")
//	}
//}
//
//func TestEvaluator(t *testing.T) {
//	if eval, err := New(""); eval != nil || err == nil {
//		t.Error("evaluator should fail for empty expression")
//	}
//
//	if eval, err := New("A|!B|  1A"); eval != nil || err == nil {
//		t.Error("evaluator should fail for unexpected character")
//	}
//
//	if eval, err := New("(A|B&(C|D)"); eval != nil || err == nil {
//		t.Error("evaluator should fail for brackets mismatch")
//	}
//
//	eval, err := New("(A|B&  C | !(Dad&Mom))")
//	if err != nil {
//		t.Error("evaluator has correct expression but failed?")
//	}
//
//	if len(eval.Parameters) != 5 {
//		t.Error("evaluator hasn't collected all parameters")
//	}
//
//	if res, err := eval.Evaluate(); err != nil || res != true {
//		t.Error("evalute doesn't worked correct")
//	}
//
//	eval.Parameters["Dad"] = true
//	eval.Parameters["Mom"] = true
//
//	if res, _ := eval.Evaluate(); res != false {
//		t.Error("arguments doesn't affected the evaluation")
//	}
//
//	eval, err = New("(A|B(C|D))")
//	if eval == nil || err != nil {
//		t.Error("evaluator shouldn't fail")
//	}
//
//	if _, err := eval.Evaluate(); err == nil {
//		t.Error("evaluate should return error for missing operator")
//	}
//
//	eval, _ = New("A&!B|!C&(G|D&(Mom|Dad))")
//	if res, err := eval.Evaluate(); err != nil || res != false {
//		t.Error("evaluate doesn't worked correct")
//	}
//}
