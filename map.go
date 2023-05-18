package suggestionmap

import (
	"os"
	"regexp"

	"gopkg.in/yaml.v3"

	_ "embed"
)

type RegularError struct {
	MatchRegex         string `json:"match,omitempty" yaml:"match"`
	ReplacementPattern string `json:"replacement,omitempty" yaml:"replacement"`
}

type SuggestionBox struct {
	matchers map[string][]RegularError
}

//go:embed default.yaml
var defaultYamlFile []byte

func LoadFile(path string) *SuggestionBox {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return load(bytes)
}

func Init() *SuggestionBox {
	return load(defaultYamlFile)
}

func load(bytes []byte) *SuggestionBox {
	var matchers map[string][]RegularError
	err := yaml.Unmarshal(bytes, &matchers)
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
