package suggestionmap

import (
	"os"
	"regexp"

	"gopkg.in/yaml.v3"
)

type RegularError struct {
	MatchRegex         string `json:"match,omitempty" yaml:"match"`
	ReplacementPattern string `json:"replacement,omitempty" yaml:"replacement"`
}

type RegularErrorMatcher struct {
	Context  string         `json:"context,omitempty" yaml:"context"`
	Matchers []RegularError `json:"matchers,omitempty" yaml:"matchers"`
}

type SuggestionBox struct {
	matchers map[string][]RegularError
}

func LoadFile(path string) *SuggestionBox {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var matchers map[string][]RegularError
	err = yaml.Unmarshal(bytes, &matchers)
	if err != nil {
		panic(err)
	}
	return &SuggestionBox{
		matchers: matchers,
	}
}

func (sb *SuggestionBox) FindSuggestions(context string, input string) []string {
	tests := sb.matchers[context]
	if len(tests) == 0 {
		return nil
	}

	suggestions := make([]string, 0)
	for _, test := range tests {
		re := regexp.MustCompile(test.MatchRegex)
		if !re.MatchString(input) {
			continue
		}

		suggestion := re.ReplaceAllString(input, test.ReplacementPattern)
		if suggestion == "" {
			continue
		}
		suggestions = append(suggestions, suggestion)
	}

	return suggestions
}
