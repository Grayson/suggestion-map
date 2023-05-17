package suggestionmap

import "testing"

func TestFindSuggestions(t *testing.T) {
	type args struct {
		context string
		input   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Basic clang-style error message",
			args{
				"clang",
				"path/to/example.swift:19:9 Cannot find 'foo' in scope",
			},
			"Make sure that foo is defined above line 19 in example.swift",
		},
		{
			"Basic go-style error message",
			args{
				"go",
				"./map.go:17:2: undefined: x",
			},
			"Make sure that x is defined above line 17 in map.go",
		},
		{
			"Empty context",
			args{
				"",
				"./map.go:17:2: undefined: x",
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindSuggestions(tt.args.context, tt.args.input); got != tt.want {
				t.Errorf("FindSuggestions() = %v, want %v", got, tt.want)
			}
		})
	}
}
