package suggestionmap

import (
	"regexp"
)

type regularError struct {
	matchRegex         string
	replacementPattern string
}

var matchers = map[string][]regularError{
	"clang": {
		{
			`^(.*):(\d+):(\d+) Cannot find '(.+?)' in scope$`,
			"Make sure that $4 is defined above line $2 in $1",
		},
	},
	"go": {
		{
			`^(.*):(\d+):(\d+): undefined: (.+)$`,
			"Make sure that $4 is defined above line $2 in $1",
		},
		{
			`^(.*):(\d+):(\d+): no new variables on left side of :=$`,
			"Looks like you're reusing a variable that's already been declared on line $2 of file $1",
		},
		{
			`^(.*):(\d+):(\d+): cannot use 1 (.+?) as (.+?) value in assignment$`,
			"Looks like a variable was defined earlier as $4 but you're setting it to $5 on line $2 of file $1.",
		},
	},
}

func FindSuggestions(context string, input string) []string {
	tests := matchers[context]
	if len(tests) == 0 {
		return nil
	}

	suggestions := make([]string, 0)
	for _, test := range tests {
		re := regexp.MustCompile(test.matchRegex)
		if !re.MatchString(input) {
			continue
		}

		suggestion := re.ReplaceAllString(input, test.replacementPattern)
		if suggestion == "" {
			continue
		}
		suggestions = append(suggestions, suggestion)
	}

	return suggestions
}
