package valueprocessing

import (
	"sort"
)

func Interpolate(data map[string]string) map[string]string {
	result := make(map[string]string)
	for _, v := range data {
		// if v starts with " or ' nothings happens
		//Otherwise:
		// should match v with : somethingORnothing ${x} somethingORnothing

		// Extract all x and substitute valid keys

		res := map[pair]string{}
		res = handleval(v, res)
		keys := make([]pair, 0, len(res))
		for k := range res {
			keys = append(keys, k)
		}
		sort.Slice(keys, func(i, j int) bool {
			return keys[i].start < keys[j].end
		})

	}
	return result
}

func handleval(val string, res map[pair]string) map[pair]string {
	stack := []int{} // Stores indexes of {
	for i, c := range val {
		if c == '{' {
			if i != 0 {
				if val[i-1] == '$' || len(stack) != 0 {
					stack = append(stack, i) // opening brackets
				}
			}
		}

		if c == '}' {
			if len(stack) != 0 {
				open := stack[len(stack)-1]  // index of the opening bracket
				stack = stack[:len(stack)-1] // pop it
				inner := val[open+1 : i]
				res[pair{
					start: open + 1,
					end:   i - 1,
				}] = inner
			}
		}
	}
	return res
}

type pair struct {
	start int
	end   int
}
