package main

import (
	"regexp"
	"strconv"

	"github.com/RosyGraph/tally/stack"
)

var operators = map[string]interface{}{
	"+": nil,
	"-": nil,
	"*": nil,
	"/": nil,
}

func ParseLaTeX(s string) string {
	r := regexp.MustCompile(`\s`)

	nums := stack.FloatStack{}
	ops := stack.StringStack{}
	tokens := r.Split(s, -1)
	for _, t := range tokens {
		if _, ok := operators[t]; ok {
			ops.Add(t)
			continue
		}

		n, err := strconv.ParseFloat(t, 32)
		if err != nil {
			panic(err)
		}

		numStack.Add(float32(n))
	}
	return tokens.String()
}
