package generator

//----------------------------
// generated code
//----------------------------

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
