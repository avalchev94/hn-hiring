package boolean

import (
	"fmt"
)

type reader struct {
	data      []byte
	readIndex int
}

func newReader(expression string) *reader {
	return &reader{
		data:      []byte(expression),
		readIndex: 0,
	}
}

func (r *reader) len() int {
	return len(r.data) - r.readIndex
}

func (r *reader) clear(char rune) error {
	for r.len() > 0 {
		ch, _ := r.read()

		if ch != char {
			r.unread()
			break
		}
	}

	return nil
}

func (r *reader) read() (rune, error) {
	if r.len() <= 0 {
		return -1, fmt.Errorf("reader's data has finished")
	}

	defer func() { r.readIndex++ }()
	return rune(r.data[r.readIndex]), nil
}

func (r *reader) seek() (rune, error) {
	if r.len() <= 0 {
		return -1, fmt.Errorf("reader's data has finished")
	}

	return rune(r.data[r.readIndex]), nil
}

func (r *reader) unread() {
	r.readIndex--
}

func (r *reader) currentIndex() int {
	return r.readIndex
}

func (r *reader) readOperator() (operator, error) {
	ch, err := r.read()
	if err != nil {
		return operator{0, -1, -1}, fmt.Errorf("failed to read rune")
	}

	switch ch {
	case or.char:
		return or, nil
	case and.char:
		return and, nil
	case not.char:
		return not, nil
	case leftBracket.char:
		return leftBracket, nil
	case rightBracket.char:
		return rightBracket, nil
	}

	r.unread()
	return operator{0, -1, -1}, fmt.Errorf("not an operator")
}

func (r *reader) readKeyword() (string, error) {
	ch, err := r.seek()
	if err != nil {
		return "", fmt.Errorf("failed to read rune")
	}

	var startQuote rune
	if ch == '"' || ch == '\'' {
		startQuote, _ = r.read()
	} else {
		return "", fmt.Errorf("keyword should start with quote")
	}

	var keyword string
	for r.len() > 0 {
		ch, _ := r.read()
		if ch == startQuote {
			break
		}
		keyword += string(ch)
	}

	if len(keyword) == 0 {
		return "", fmt.Errorf("quotes are emtpy filled")
	}

	return keyword, nil
}
