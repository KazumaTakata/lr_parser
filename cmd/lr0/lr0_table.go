package lr0

import (
	"fmt"
	"github.com/KazumaTakata/ascii_graph"
	"github.com/KazumaTakata/lr_parser/util"
	"strconv"
	//	"strings"
)

type Action_type int

const (
	Shift  Action_type = 0
	Accept Action_type = 1
	Reduce Action_type = 2
)

func (action Action_type) String() string {
	actions := []string{
		"Shift",
		"Accrpt",
		"Reduce"}

	return actions[action]
}

type Reduction struct {
	left  string
	right []string
}

type Action struct {
	action_type Action_type
	reduction   Reduction
}

type Table_element struct {
	goto_table map[string]int
	action     Action
}

type Table struct {
	table_elements []Table_element
}

func Construct_lr0_Table(state_with_next_list []State_with_next, bnf_list []util.Bnf) Table {

	table := Table{}

	for _, state_with_next := range state_with_next_list {
		table_element := Table_element{goto_table: state_with_next.Next}

		handlers := Get_handlers(state_with_next, bnf_list)
		if len(handlers) > 0 {
			if len(handlers) > 1 {
				fmt.Printf("reduction conflicts")
			}
			handler := handlers[0]
			reduction := Reduction{left: bnf_list[handler.Product_id].Left, right: bnf_list[handler.Product_id].Right[handler.Alternate_id]}
			action := Action{action_type: Reduce, reduction: reduction}

			table_element.action = action
		} else {
			action := Action{action_type: Shift}
			table_element.action = action
		}

		table.table_elements = append(table.table_elements, table_element)
	}

	return table
}

func (table *Table) Print_lr0_table(bnf_list []util.Bnf) {

	ascii_table := ascii_graph.Table2d{}
	col_p := []string{}

	for i := 0; i < len(table.table_elements); i++ {
		col_p = append(col_p, strconv.Itoa(i))
	}
	nonterminal_and_terminal := util.Get_nonterminal_and_terminal(bnf_list)
	row_p := []string{}
	for term, _ := range nonterminal_and_terminal {
		row_p = append(row_p, term)
	}
	data := [][]string{}
	for i, table_ele := range table.table_elements {
		data = append(data, []string{})
		for _, term := range row_p {
			if index, ok := table_ele.goto_table[term]; ok {
				data[i] = append(data[i], strconv.Itoa(index))
			} else {
				data[i] = append(data[i], " ")
			}

		}
	}

	row_p = append(row_p, "Action")
	for i, table_ele := range table.table_elements {
		action := table_ele.action.action_type.String()
		if len(table_ele.action.reduction.left) > 0 {
			action = action + " "

			action = action + table_ele.action.reduction.left
			action = action + "->"
			for _, red := range table_ele.action.reduction.right {
				action = action + red
			}
		}
		data[i] = append(data[i], action)
	}
	ascii_table.SetColumnProperty(col_p)
	ascii_table.SetRowProperty(row_p)
	ascii_table.SetData(data)

	ascii_table.Draw()

}
