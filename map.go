package suggestionmap

import (
	"fmt"
	"path"
	"regexp"
)

const (
	test = `(.*):(\d+):(\d+) Cannot find '(.+?)' in scope`
)

func FindSuggestions(context string, input string) string {
	re := regexp.MustCompile(test)
	matches := re.FindStringSubmatch(input)
	return fmt.Sprintf("Make sure that %v is defined above line %v in %v", matches[4], matches[2], path.Base(matches[1]))
}
