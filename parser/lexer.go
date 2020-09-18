package parser

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

func (l *Lexer) scan() (*Word, error) {
	s := l.source[l.now]
	l.now++
	return getWord(s)
}

func (l *Lexer) str() string {
	return string(l.source)
}
