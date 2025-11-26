package parser

import (
	"errors"
	"fmt"
	"strings"

	"github.com/indexdata/ccms/cmd/ccd/ast"
)

//go:generate ragel -Z -G2 -o scan.go scan.rl
//go:generate go tool goyacc -l -o gram.go gram.y

func WriteErrorContext(query string, ts, te int) string {
	var b strings.Builder
	// Scan to position, counting the number of newlines.
	var pos, line, markline, linepos int
	for _, r := range query {
		if pos >= ts {
			markline = line
			break
		}
		if r == '\n' {
			line++
			linepos = 0
			pos++
			continue
		}
		pos++
		linepos++
	}
	//s := fmt.Sprintf("line %d: ", markline+1)
	//margin := len(s)
	//b.WriteString(s)
	margin := 0
	// Scan again, printing the line containing position.
	var w bool
	pos = 0
	line = 0
	for _, r := range query {
		if line >= markline {
			if r != '\n' {
				b.WriteRune(r)
			}
			w = true
		}
		if r == '\n' {
			if w {
				break
			}
			line++
		}
		pos++
	}
	b.WriteRune('\n')
	// Write pointer at linepos.
	for i := 0; i < margin; i++ {
		b.WriteRune(' ')
	}
	for i := 0; i < linepos; i++ {
		b.WriteRune(' ')
	}
	WriteCarets(&b, ts, te)
	return b.String()
}

func errorMessage(l *lexer) error {
	ts := l.ts
	te := l.te
	if ts == te {
		te++
	}
	s := fmt.Sprintf("%s near %q\n%s", l.err, l.data[ts:te], WriteErrorContext(string(l.data), ts, te))
	return errors.New(s)
}

func Parse(input string) (ast.Node, error, bool) {
	l := newLexer([]byte(input))
	e := yyParse(l)
	var msg error
	if e == 0 {
		if l.node == nil {
			var b strings.Builder
			WriteCarets(&b, l.ts, l.te)
			msg = fmt.Errorf("syntax error near %q\n%s\n%s",
				strings.Split(input, " ")[0], strings.Split(input, "\n")[0], b.String())
		}
	} else {
		msg = errorMessage(l)
	}
	return l.node, msg, l.pass
}

func WriteCarets(b *strings.Builder, ts, te int) {
	for i := ts; i < te; i++ {
		b.WriteRune('^')
	}
}
