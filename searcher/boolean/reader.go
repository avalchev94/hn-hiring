package boolean

import (
	"errors"
)

type reader struct {
	stream    []byte
	readIndex int
}

func newReader(expression string) *reader {
	return &reader{
		stream:    []byte(expression),
		readIndex: 0,
	}
}

func (r *reader) len() int {
	return len(r.stream) - r.readIndex
}

func (r *reader) clear(char rune) error {
	for r.len() > 0 {
		ch, _ := r.read()

		if ch != char {
			r.readIndex--
			break
		}
	}

	return nil
}

func (r *reader) read() (rune, error) {
	if r.len() <= 0 {
		return -1, errors.New("end of stream")
	}

	defer func() { r.readIndex++ }()
	return rune(r.stream[r.readIndex]), nil
}

func (r *reader) currentIndex() int {
	return r.readIndex
}

func (r *reader) readToken() (interface{}, error) {
	ch, err := r.read()
	if err != nil {
		return nil, errors.New("failed to read")
	}

	switch ch {
	// keyword
	case '"':
		keyword := ""
		for r.len() > 0 {
			ch, _ = r.read()
			if ch == '"' {
				if keyword != "" {
					return keyword, nil
				}
				return nil, errors.New("empty keyword")
			}
			keyword += string(ch)
		}
		return nil, errors.New("quotes does not match")
	// operators
	case and.char:
		return and, nil
	case or.char:
		return or, nil
	case not.char:
		return not, nil
	case leftBracket.char:
		return leftBracket, nil
	case rightBracket.char:
		return rightBracket, nil
	}

	return nil, errors.New("unexpected character")
}
