package suggestionmap

import (
	"fmt"
	"path"
	"regexp"
)

var matchers = map[string]string{
	"clang": `(.*):(\d+):(\d+) Cannot find '(.+?)' in scope`,
	"go":    `(.*):(\d+):(\d+): undefined: (.+)`,
}

func FindSuggestions(context string, input string) string {
	test := matchers[context]
	if test == "" {
		return ""
	}

	re := regexp.MustCompile(test)
	matches := re.FindStringSubmatch(input)
	if len(matches) == 0 {
		return ""
	}

	return fmt.Sprintf("Make sure that %v is defined above line %v in %v", matches[4], matches[2], path.Base(matches[1]))
}
