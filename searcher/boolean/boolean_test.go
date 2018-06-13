package boolean

import (
//	"testing"

//	"github.com/avalchev94/go_collections/stack"
)

//
//func TestReader(t *testing.T) {
//	r := newReader("  'A'|'Bac12'&  !('1'|'A')")
//
//	if r == nil {
//		t.Errorf("newReader shouldn't be nil")
//	}
//
//	if r.len() != 18 {
//		t.Errorf("Length should be 18, nothing is read currently")
//	}
//
//	if err := r.clear(' '); err != nil {
//		t.Errorf("Clear shouldn't had to return nil")
//	}
//
//	if r.len() != 16 {
//		t.Errorf("After clearing the whitespaces, len should be 16")
//	}
//
//	if ch, err := r.read(); err != nil {
//		t.Errorf("Read shouldn't return error")
//	} else if ch != 'A' {
//		t.Errorf("The read character is not A")
//	}
//
//	if r.len() != 15 {
//		t.Errorf("After reading the first rune, len should be 15")
//	}
//
//	if ch, err := r.seek(); err != nil {
//		t.Errorf("Seek shouldn't return error")
//	} else if ch != '|' {
//		t.Errorf("Seek returned wrong character")
//	}
//
//	if r.len() != 15 {
//		t.Errorf("Seek should only check the char, but not mark it as read")
//	}
//
//	if op, err := r.readOperator(); err != nil {
//		t.Errorf("read operator should be successful")
//	} else if !op.equal(or) {
//		t.Errorf("wrong operator returned")
//	}
//
//	if k, err := r.readKeyword(); err != nil {
//		t.Errorf("read parameter should be successful")
//	} else if k != "Bac12" {
//		t.Errorf("wrong parameter returned")
//	}
//
//	if op, _ := r.readOperator(); !op.equal(and) {
//		t.Errorf("wrong operator returned")
//	}
//
//	r.clear(' ')
//	if op, _ := r.readOperator(); !op.equal(not) {
//		t.Errorf("wrong operator returned")
//	}
//
//	if op, _ := r.readOperator(); !op.equal(leftBracket) {
//		t.Errorf("wrong operator returned")
//	}
//
//	if _, err := r.readParameter(); err == nil {
//		t.Errorf("error should be returned")
//	}
//
//	if _, err := r.readOperator(); err == nil {
//		t.Errorf("error should be returned")
//	}
//
//	r.read()
//	r.readOperator()
//	r.readParameter()
//
//	if op, _ := r.readOperator(); !op.equal(rightBracket) {
//		t.Errorf("wrong operator returned")
//	}
//
//	if r.len() != 0 {
//		t.Errorf("data has ended. 0 should be returned")
//	}
//
//	_, err0 := r.read()
//	_, err1 := r.seek()
//	_, err2 := r.readOperator()
//	_, err3 := r.readParameter()
//
//	if err0 == nil || err1 == nil || err2 == nil || err3 == nil {
//		t.Errorf("data has ended, error should be returned")
//	}
//}
//
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
