package lr0

import (
	"io/ioutil"
	//	"remove_left_recursion/ll"
	"github.com/KazumaTakata/lr_parser/cmd/util"
	"github.com/KazumaTakata/lr_parser/util"
	//	"strings"
)

type State_with_next struct {
	Next  map[string]int
	State u_state.State
}

func Lr0_automata(filepath string) ([]State_with_next, []util.Bnf) {

	bnf, err := ioutil.ReadFile(filepath)
	util.Check(err)

	bnf_parsed := util.Parse_bnf_file(string(bnf))

	start_state := gen_start_state(bnf_parsed)

	automaton_states := []State_with_next{}
	_, root_index := add_to_automaton_states(&automaton_states, start_state)
	not_explored_queue := u_state.Not_explored_queue{Queue: []int{root_index}}

	for !not_explored_queue.Empty() {
		root_index := not_explored_queue.Dequeue()
		root_state := automaton_states[root_index]
		new_states := create_new_states(bnf_parsed, root_state.State)
		not_explored := add_all_to_automaton_states(&automaton_states, root_index, new_states)

		for _, index := range not_explored {
			not_explored_queue.Enqueue(index)
		}
	}

	return automaton_states, bnf_parsed
}

func is_equal(state_a, state_b u_state.State) bool {
	if len(state_a) != len(state_b) {
		return false
	}

	for state_a_ele, _ := range state_a {
		if _, ok := state_b[state_a_ele]; !ok {
			return false
		}
	}

	return true
}

func add_to_automaton_states(automaton_state *[]State_with_next, new_state u_state.State) (bool, int) {

	for index, state := range *automaton_state {
		if is_equal(state.State, new_state) {
			return false, index
		}
	}

	*automaton_state = append(*automaton_state, State_with_next{State: new_state, Next: make(map[string]int)})

	return true, len(*automaton_state) - 1
}

func gen_start_state(bnf_list []util.Bnf) u_state.State {
	start_state := make(u_state.State)
	start_state[u_state.State_element{Product_id: 0, Alternate_id: 0, Offset: 0}] = true

	new_start_state := make(u_state.State)
	for state_ele, _ := range start_state {
		new_elements := Expand_non_terminal(state_ele, bnf_list, u_state.State{})
		for new_ele, _ := range new_elements {
			new_start_state[new_ele] = true
		}
	}
	for new_elem, _ := range new_start_state {
		start_state[new_elem] = true
	}

	return start_state

}

func Expand_non_terminal(state_element u_state.State_element, bnf_list []util.Bnf, current_state u_state.State) u_state.State {

	non_terminals := util.Get_nonterminal(bnf_list)
	added_state := make(u_state.State)

	if state_element.Offset < len(bnf_list[state_element.Product_id].Right[state_element.Alternate_id]) {
		right_ele := bnf_list[state_element.Product_id].Right[state_element.Alternate_id][state_element.Offset]
		if _, ok := non_terminals[right_ele]; ok {
			for prod_id, bnf := range bnf_list {
				if bnf.Left == right_ele {
					for alte_id, _ := range bnf.Right {
						new_state_element := u_state.State_element{Product_id: prod_id, Alternate_id: alte_id, Offset: 0}

						if _, ok := current_state[new_state_element]; ok {
							return u_state.State{}
						}

						added_state[new_state_element] = true
						current_state[new_state_element] = true

						new_state := Expand_non_terminal(new_state_element, bnf_list, current_state)
						for new_ele, _ := range new_state {
							added_state[new_ele] = true
						}
					}
				}
			}
		}

	}

	return added_state
}

func add_all_to_automaton_states(automaton_states *[]State_with_next, root_index int, new_states map[string]u_state.State) []int {

	not_explored := []int{}
	for key, new_state := range new_states {
		is_new, index := add_to_automaton_states(automaton_states, new_state)
		if is_new {
			not_explored = append(not_explored, index)
		}
		(*automaton_states)[root_index].Next[key] = index
	}

	return not_explored
}

func create_new_states(bnf_list []util.Bnf, root_state u_state.State) map[string]u_state.State {

	nonterminal_and_terminal := util.Get_nonterminal_and_terminal(bnf_list)

	new_state_elements := map[string][]u_state.State_element{}

	for node, _ := range nonterminal_and_terminal {
		for state_ele, _ := range root_state {
			if !u_state.Is_last(state_ele, bnf_list) {
				if node == bnf_list[state_ele.Product_id].Right[state_ele.Alternate_id][state_ele.Offset] {
					new_state_elements[node] = append(new_state_elements[node], u_state.State_element{Product_id: state_ele.Product_id, Alternate_id: state_ele.Alternate_id, Offset: state_ele.Offset + 1})
				}
			}
		}

	}
	new_states := map[string]u_state.State{}

	for key, elements := range new_state_elements {
		new_state := u_state.State{}
		for _, element := range elements {
			new_elements := Expand_non_terminal(element, bnf_list, u_state.State{})
			for new_ele, _ := range new_elements {
				new_state[new_ele] = true
			}
			new_state[element] = true
		}
		new_states[key] = new_state
	}

	return new_states
}
