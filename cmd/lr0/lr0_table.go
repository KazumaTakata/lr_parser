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
	Left  string
	Right []string
}

type Action struct {
	Action_type Action_type
	Reduction   Reduction
}

type Table_element struct {
	Goto_table map[string]int
	Action     Action
}

type Table struct {
	Table_elements []Table_element
}

func Construct_lr0_Table(state_with_next_list []State_with_next, bnf_list []util.Bnf) Table {

	table := Table{}

	for _, state_with_next := range state_with_next_list {
		table_element := Table_element{Goto_table: state_with_next.Next}

		handlers := Get_handlers(state_with_next, bnf_list)
		if len(handlers) > 0 {
			if len(handlers) > 1 {
				fmt.Printf("reduction conflicts")
			}
			handler := handlers[0]
			reduction := Reduction{Left: bnf_list[handler.Product_id].Left, Right: bnf_list[handler.Product_id].Right[handler.Alternate_id]}
			action := Action{Action_type: Reduce, Reduction: reduction}

			table_element.Action = action
		} else {
			action := Action{Action_type: Shift}
			table_element.Action = action
		}

		table.Table_elements = append(table.Table_elements, table_element)
	}

	return table
}

func (table *Table) Print_lr0_table(bnf_list []util.Bnf) {

	ascii_table := ascii_graph.Table2d{}
	col_p := []string{}

	for i := 0; i < len(table.Table_elements); i++ {
		col_p = append(col_p, strconv.Itoa(i))
	}
	nonterminal_and_terminal := util.Get_nonterminal_and_terminal(bnf_list)
	row_p := []string{}
	for term, _ := range nonterminal_and_terminal {
		row_p = append(row_p, term)
	}
	data := [][]string{}
	for i, table_ele := range table.Table_elements {
		data = append(data, []string{})
		for _, term := range row_p {
			if index, ok := table_ele.Goto_table[term]; ok {
				data[i] = append(data[i], strconv.Itoa(index))
			} else {
				data[i] = append(data[i], " ")
			}

		}
	}

	row_p = append(row_p, "Action")
	for i, table_ele := range table.Table_elements {
		action := table_ele.Action.Action_type.String()
		if len(table_ele.Action.Reduction.Left) > 0 {
			action = action + " "

			action = action + table_ele.Action.Reduction.Left
			action = action + "->"
			for _, red := range table_ele.Action.Reduction.Right {
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
