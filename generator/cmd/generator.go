package main

import (
	"io/ioutil"
	"os"
	"text/template"
)

type Package struct {
	Name      string
	Functions []Function
	Structs   []Struct
}

type Function struct {
	Name       string
	Reciever   string
	Statements []Statement
}

type Statement struct {
	AssignStatement string
	IfStatement     *IfStatement
	FunctionCall    *FunctionCall
}

type AssignStatement struct {
	Left  string
	Right *Expression
}

type Expression struct {
}

type IfStatement struct {
	BoolExpression *BooleanExpression
	Body           []Statement
	Else           *IfStatement
}

type BooleanExpression struct {
	Left  string
	Right string
}

type FunctionCall struct {
	Name string
	Args []string
}

type Struct struct {
	Name     string
	Elements []StructElement
}

type StructElement struct {
	Name string
	Type string
}

func main() {

	package_temp := Package{Name: "main", Functions: []Function{}}
	mainfunc_statement := Statement{AssignStatement: "i := 0"}
	mainfunc_ifstatement := Statement{IfStatement: &IfStatement{BoolExpression: &BooleanExpression{Left: "i", Right: "0"}, Body: []Statement{mainfunc_statement}}}
	mainfunc := Function{Name: "main", Reciever: "(s S)", Statements: []Statement{mainfunc_statement, mainfunc_ifstatement}}
	package_temp.Functions = append(package_temp.Functions, mainfunc)
	new_struct := Struct{Name: "Node", Elements: []StructElement{StructElement{Name: "Id", Type: "string"}, StructElement{Name: "T", Type: "*T"}}}
	package_temp.Structs = append(package_temp.Structs, new_struct)

	buf, _ := ioutil.ReadFile("../sample.tmpl")

	tmpl, err := template.New("package").Parse(string(buf))
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, package_temp)
	if err != nil {
		panic(err)
	}

}
