package main

import "fmt"

func main() {
	result := myAtoi("    -42")
	fmt.Printf("result: %v\n", result)
}

func myAtoi(s string) int {
	// start signed number end
	const MinInt32, MaxInt32 = -1 << 31, 1<<31 - 1
	states := map[string][]string{
		"start":  {"end", "signed", "number", "end"},
		"signed": {"end", "end", "number", "end"},
		"number": {"end", "end", "number", "end"},
		"end":    {"end", "end", "end", "end"},
	}

	index := map[string]int{
		"start":  0,
		"signed": 1,
		"number": 2,
		"end":    3,
	}

	state := "start"
	ans := 0
	sign := 1

	for i := 0; i <= len(s); i++ {
		curr := get_state(s, i)

		switch curr {
		case "end":
			fmt.Printf("ans: %v\n", ans)
			fmt.Printf("sign: %v\n", sign)
			return ans * sign
		case "number":
			ans = ans*10 + int(s[i]) - 48
			if ans >= MaxInt32 {
				return MaxInt32
			}
			if ans <= MinInt32 {
				return MinInt32
			}
		case "signed":
			if s[i] == '-' {
				sign = -1
			}
		}

		curr = states[state][index[curr]]
	}

	return ans
}

func get_state(s string, i int) string {
	switch {
	case i >= len(s):
		return "end"
	case s[i] == ' ':
		return "start"
	case s[i] == '+' || s[i] == '-':
		return "signed"
	case s[i] >= '0' && s[i] <= '9':
		return "number"
	}
	return "end"
}
