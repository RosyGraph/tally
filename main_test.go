package main

import "testing"

func TestTally(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "single int",
			s:    "1",
			want: "1",
		},
		{
			name: "add two ints",
			s:    "2 + 2",
			want: "4",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ParseLaTeX(test.s)
			if got != test.want {
				t.Errorf("[%s]\ngot: %s want: %s", test.name, got, test.want)
			}
		})
	}
}
