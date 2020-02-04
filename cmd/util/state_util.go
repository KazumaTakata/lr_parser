package u_state

import (
	"github.com/KazumaTakata/lr_parser/util"
	//	"strings"
)

type Not_explored_queue struct {
	Queue []int
}

func (q *Not_explored_queue) Enqueue(new_item int) {
	q.Queue = append(q.Queue, new_item)
}

func (q *Not_explored_queue) Dequeue() int {
	dequeued := q.Queue[0]
	q.Queue = q.Queue[1:]
	return dequeued
}

func (q *Not_explored_queue) Empty() bool {
	if len(q.Queue) == 0 {
		return true
	}

	return false
}

func Is_last(state_element State_element, bnf_list []util.Bnf) bool {
	if len(bnf_list[state_element.Product_id].Right[state_element.Alternate_id]) > state_element.Offset {
		return false
	}

	return true
}

type State map[State_element]bool

type State_element struct {
	Product_id   int
	Alternate_id int
	Offset       int
}
