package catalog

import (
	"strings"
)

// given a project name or fully qualified set name (project.set), return true if the project exists
func ProjectExists(projectOrFullSetName string) bool {
	var p string // project name
	sp := strings.Split(projectOrFullSetName, ".")
	switch len(sp) {
	case 1: // project
		fallthrough
	case 2: // full set name
		p = sp[0]
	}
	if p == "test" {
		return true
	}
	return false
}
