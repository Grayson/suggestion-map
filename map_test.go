package suggestionmap

import "testing"

const (
	multilineGoError = `./map.go:20:7: no new variables on left side of :=
./map.go:20:10: cannot use 1 (untyped int constant) as string value in assignment`
)

func TestFindSuggestions(t *testing.T) {
	type args struct {
		context string
		input   string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"Basic clang-style error message",
			args{
				"clang",
				"path/to/example.swift:19:9 Cannot find 'foo' in scope",
			},
			[]string{"Make sure that foo is defined above line 19 in path/to/example.swift"},
		},
		{
			"Basic go-style error message",
			args{
				"go",
				"./map.go:17:2: undefined: x",
			},
			[]string{"Make sure that x is defined above line 17 in ./map.go"},
		},
		{
			"Empty context",
			args{
				"",
				"./map.go:17:2: undefined: x",
			},
			[]string{""},
		},
		{
			"Empty input",
			args{
				"go",
				"",
			},
			[]string{""},
		},
		{
			"Multiline Go error",
			args{
				"go",
				multilineGoError,
			},
			[]string{
				"Looks like you're reusing a variable that's already been declared on line 20 of file ./map.go",
				"Looks like a variable was defined earlier as untyped int constant but you're setting it to string on line 20 of file ./map.go.",
			},
		},
	}

	sb := Init()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sb.FindSuggestions(tt.args.context, tt.args.input)
			for idx := 0; idx < len(got); idx++ {
				if got[idx] == tt.want[idx] {
					continue
				}
				t.Errorf("FindSuggestions() = %v, want %v", got[idx], tt.want[idx])
			}
		})
	}
}
