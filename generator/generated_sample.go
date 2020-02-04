package generator

import "fmt"

//----------------------------
// generated code
//----------------------------

func PrintTree(node ParserNode, indent string) {

	if s_node, ok := node.(S); ok {
		fmt.Printf(indent)
		fmt.Printf("|-")
		fmt.Println("S")
		if s_node.E != nil {
			indent = indent + "| "
			PrintTree(*s_node.E, indent)
		}
	}

	if s_node, ok := node.(E); ok {

		fmt.Printf(indent)
		fmt.Printf("|-")
		fmt.Println("E")
		if s_node.T != nil {
			PrintTree(*s_node.T, indent+"| ")
		}
		if s_node.E != nil {
			PrintTree(*s_node.E, indent+"| ")
		}
	}

	if s_node, ok := node.(T); ok {
		fmt.Printf(indent)
		fmt.Printf("|-")
		fmt.Println("T")

		if s_node.E != nil {
			PrintTree(*s_node.E, indent+"| ")
		}
		if s_node.Int != nil {
			PrintTree(*s_node.Int, indent+"| ")
		}
	}

	if s_node, ok := node.(Terminal); ok {
		fmt.Printf(indent)
		fmt.Printf("|-")
		fmt.Println("Terminal")

		if s_node.Value != "" {
			indent = indent + "| "
			fmt.Printf(indent)
			fmt.Printf("|-")
			fmt.Println(s_node.Value)
		}
	}

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
func ContructParserNode(root ParserNode, right string, right_node ParserNode) ParserNode {

	s_root := root

	if root.String() == "S" {
		if right == "E" {
			s_root, _ := root.(S)
			right_node, _ := right_node.(E)
			s_root.E = &right_node
			return s_root
		}

	} else if root.String() == "T" {
		if right == "E" {
			s_root, _ := root.(T)
			right_node, _ := right_node.(E)
			s_root.E = &right_node
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
			right_node, _ := right_node.(E)
			s_root.E = &right_node
			return s_root
		}
		if right == "T" {
			s_root, _ := root.(E)
			right_node, _ := right_node.(T)
			s_root.T = &right_node
			return s_root
		}
	}

	return s_root
}

func ContructRootNode(left string) ParserNode {

	if left == "S" {
		return S{}
	} else if left == "E" {
		return E{}
	} else if left == "T" {
		return T{}
	}

	return nil
}
