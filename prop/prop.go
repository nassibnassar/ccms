// primitives for parsing composite properties
package prop

import "strings"

// a name-title pair
type Prop struct {
	Name  string // unique identifier
	Title string // title or brief description
}

// parse the value of a composite property into a list of name-title pairs
func Parse(property string) []Prop {
	prop := make([]Prop, 0)
	pairs := strings.Split(property, "|")
	for i := range pairs {
		p := strings.Split(pairs[i], ":")
		l := len(p)
		var name, title string
		if l > 0 {
			name = p[0]
			if l > 1 {
				title = p[1]
			}
		}
		prop = append(prop, Prop{Name: name, Title: title})
	}
	return prop
}

// possibly not needed
// a list of name-title pairs
// type Property []Prop

// possibly not needed
// assemble a list of name-title pairs into a composite property value
// func (p Property) String() string {
// 	var b strings.Builder
// 	for i := range p {
// 		if i != 0 {
// 			b.WriteRune('|')
// 		}
// 		b.WriteString(p[i].Name)
// 		b.WriteRune(':')
// 		b.WriteString(p[i].Title)
// 	}
// 	return b.String()
// }
