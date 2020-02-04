package main

import (
	"github.com/KazumaTakata/lr_parser/cmd/lr0"
	"github.com/KazumaTakata/lr_parser/generator"
	"path/filepath"
	"runtime"
	//	"strings"
)

func main() {

	_, filename, _, _ := runtime.Caller(0)
	bnf_path := filepath.Join(filepath.Dir(filepath.Dir(filename)), "sample2.bnf")

	automaton_states, bnf_list := lr0.Lr0_automata(bnf_path)

	table := lr0.Construct_lr0_Table(automaton_states, bnf_list)

	table.Print_lr0_table(bnf_list)
	start := generator.Parse_lr0(table, []string{"Int", "+", "Int", ";"}, bnf_list)
	generator.PrintTree(start, "")

}
