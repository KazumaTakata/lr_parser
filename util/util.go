package util

import "strings"

type Bnf struct {
	Left  string
	Right [][]string
}

func parse_bnf_right(right string) [][]string {
	rights := strings.Split(right, "|")
	parsed_list := [][]string{}
	for _, right := range rights {
		parsed_list = append(parsed_list, strings.Split(right, " "))
	}

	trimed_list := [][]string{}
	for _, right := range parsed_list {
		tmp := []string{}
		for _, r := range right {

			trimed := strings.TrimSpace(r)
			if len(trimed) > 0 {
				tmp = append(tmp, trimed)
			}
		}

		trimed_list = append(trimed_list, tmp)
	}

	return trimed_list
}

func Parse_bnf_file(bnf_string string) []Bnf {

	bnf_lines := strings.Split(bnf_string, "\n")

	bnf_parsed := []Bnf{}

	for _, line := range bnf_lines {
		left_right := strings.Split(line, "::=")

		if len(left_right) > 1 {
			trimed_list := parse_bnf_right(left_right[1])
			bnf_parsed = append(bnf_parsed, Bnf{Left: strings.TrimSpace(left_right[0]), Right: trimed_list})
		}
	}

	return bnf_parsed

}

func Get_terminal(bnf_list []Bnf, non_terminal []string) map[string]bool {

	terminal := map[string]bool{}

	for _, prod := range bnf_list {
		for _, rights := range prod.Right {
			for _, right := range rights {
				terminal[right] = true
			}
		}
	}

	return terminal

}

func Get_nonterminal(bnf_list []Bnf) map[string]bool {

	non_terminal := map[string]bool{}

	for _, prod := range bnf_list {
		non_terminal[prod.Left] = true
	}

	return non_terminal

}

func Get_nonterminal_and_terminal(bnf_list []Bnf) map[string]bool {

	nonterminal_and_terminal := map[string]bool{}

	for _, prod := range bnf_list {
		for _, rights := range prod.Right {
			for _, right := range rights {
				nonterminal_and_terminal[right] = true
			}
		}

		nonterminal_and_terminal[prod.Left] = true
	}

	return nonterminal_and_terminal

}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
