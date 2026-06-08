package cat

import (
	"strings"
	"unicode"
)

var Attributes []string = []string{
	"id",
	"author",
	"title",
	"full_vendor_name",
	"availability",
}

var AttributeMap map[string]struct{}

func Init() {
	AttributeMap = make(map[string]struct{})
	for i := range Attributes {
		AttributeMap[Attributes[i]] = struct{}{}
	}
}

func IsAttr(attr string) bool {
	_, ok := AttributeMap[attr]
	return ok
}

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
