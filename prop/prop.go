package prop

import "strings"

// a name-title pair
type Prop struct {
	Name  string // unique identifier
	Title string // title or brief description
}

// parse the value of a structured property into a slice of name-title pairs
func Parse(rawProperty string) []Prop {
	prop := make([]Prop, 0)
	pairs := strings.Split(rawProperty, "|")
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
