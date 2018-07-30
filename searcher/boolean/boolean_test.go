package boolean

import (
	"testing"

	"github.com/avalchev94/go_collections/stack"
	"github.com/avalchev94/go_collections/tree"
	"github.com/stretchr/testify/assert"
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
