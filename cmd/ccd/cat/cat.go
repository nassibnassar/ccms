package cat

import (
	"strings"
	"unicode"
)

func makeTitle(name string) string {
	var b strings.Builder
	q := '_'
	for _, r := range name {
		if r == '_' {
			if q != '_' {
				b.WriteRune(' ')
			}
			q = r
			continue
		}
		if q == '_' {
			b.WriteRune(unicode.ToUpper(r))
			q = r
			continue
		}
		b.WriteRune(r)
		q = r
	}
	return b.String()
}
