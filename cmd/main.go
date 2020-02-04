package main

import (
	"fmt"
	"github.com/KazumaTakata/lr_parser/cmd/lr0"
	"github.com/KazumaTakata/lr_parser/util"
	"path/filepath"
	"runtime"
	//	"strings"
)

type ParserNode interface {
	parser_node()
	String() string
}

type S struct {
	E *E
}

func (s S) parser_node() {}

func (s S) String() string {
	return "S"
}

type E struct {
	Id int
	T  *T
	E  *E
}

func (s E) parser_node() {}

func (s E) String() string {
	return "E"
}

type T struct {
	Id  int
	E   *E
	Int *Terminal
}

func (s T) parser_node() {}

func (s T) String() string {
	return "T"
}

type Terminal struct {
	Value string
}

func (s Terminal) parser_node() {}

func (s Terminal) String() string {
	return s.Value
}

type state_stack struct {
	data []int
}

func (s *state_stack) pop() int {
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return top
}

func (s *state_stack) push(d int) {
	s.data = append(s.data, d)
}

func (s *state_stack) top() int {
	return s.data[len(s.data)-1]
}

type symbol_stack struct {
	data []ParserNode
}

func (s *symbol_stack) pop() ParserNode {
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return top
}

func (s *symbol_stack) push(d ParserNode) {
	s.data = append(s.data, d)
}

func (s *symbol_stack) top() ParserNode {
	return s.data[len(s.data)-1]
}

type node struct {
	node_type string
	children  []node
}

func parse_lr0(table lr0.Table, input_tokens []string, bnf_list []util.Bnf) symbol_stack {

	state_stack := state_stack{}
	symbol_stack := symbol_stack{}

	state_stack.push(0)

	for _, token := range input_tokens {
		symbol_stack.push(Terminal{Value: token})
		next_state_id := table.Table_elements[state_stack.top()].Goto_table[symbol_stack.top().String()]
		state_stack.push(next_state_id)
		handle_reduction(table, next_state_id, bnf_list, &symbol_stack, &state_stack)
	}

	return symbol_stack

}

func ContructParserNode(root ParserNode, right string) ParserNode {

	s_root := root

	if root.String() == "S" {
		if right == "E" {
			s_root, _ := root.(S)
			s_root.E = &E{}
			return s_root
		}

	} else if root.String() == "T" {
		if right == "E" {
			s_root, _ := root.(T)
			s_root.E = &E{}
			return s_root
		}
		if right == "Int" {
			s_root, _ := root.(T)
			s_root.Int = &Terminal{Value: "Int"}
			return s_root
		}
	} else if root.String() == "E" {
		if right == "E" {
			s_root, _ := root.(E)
			s_root.E = &E{}
			return s_root
		}
		if right == "T" {
			s_root, _ := root.(E)
			s_root.T = &T{}
			return s_root
		}
	}

	return s_root
}

func handle_reduction(table lr0.Table, next_state_id int, bnf_list []util.Bnf, symbol_stack *symbol_stack, state_stack *state_stack) {

	next_state := table.Table_elements[next_state_id]
	if next_state.Action.Action_type == lr0.Reduce {

		right := next_state.Action.Reduction.Right
		left := next_state.Action.Reduction.Left
		//root_node := node{node_type: left, children: []node{}}
		var root_node ParserNode

		if left == "S" {
			root_node = S{}
		} else if left == "E" {
			root_node = E{}
		} else if left == "T" {
			root_node = T{}
		}

		for i := len(right) - 1; i >= 0; i-- {
			poped := symbol_stack.pop()
			if poped.String() == right[i] {
				root_node = ContructParserNode(root_node, right[i])
				//root_node.children = append([]node{poped}, root_node.children...)
				state_stack.pop()
			} else {
				fmt.Printf("parse error")
			}
		}
		symbol_stack.push(root_node)
		next_state_id = table.Table_elements[state_stack.top()].Goto_table[symbol_stack.top().String()]
		state_stack.push(next_state_id)
		handle_reduction(table, next_state_id, bnf_list, symbol_stack, state_stack)
	}

}

func main() {

	_, filename, _, _ := runtime.Caller(0)
	bnf_path := filepath.Join(filepath.Dir(filepath.Dir(filename)), "sample2.bnf")

	automaton_states, bnf_list := lr0.Lr0_automata(bnf_path)

	table := lr0.Construct_lr0_Table(automaton_states, bnf_list)

	table.Print_lr0_table(bnf_list)
	stack := parse_lr0(table, []string{"Int", "+", "Int", ";"}, bnf_list)
	fmt.Printf("%+v", stack)

}
