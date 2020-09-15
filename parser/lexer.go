package parser

import "errors"

type Lexer struct {
	source []rune
	now    int
}

func newLexer(source *string) *Lexer {
	return &Lexer{
		source: []rune(*source + "$"),
		now:    0,
	}
}

func (l *Lexer) lookAt() rune {
	return l.source[l.now]
}

func (l *Lexer) expect(s rune) error {
	if l.source[l.now] == s {
		l.now++
		return nil
	} else {
		return errors.New("expect " + string(s) + ", but found " + string(l.source[l.now]))
	}
}

func (l *Lexer) scan() (*Word, error) {
	s := l.source[l.now]
	err := l.expect(s)
	if err == nil {
		return GetWord(s)
	} else {
		return nil, err
	}
}
