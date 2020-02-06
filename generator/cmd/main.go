package main

import (
	"fmt"
	"github.com/KazumaTakata/go_code_generator"
	"github.com/KazumaTakata/lr_parser/util"
	"github.com/KazumaTakata/regex_virtualmachine"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

func CreateStruct(bnf_list []util.Bnf) []generator.Struct {

	structs := []generator.Struct{}
	nonterminals := util.Get_nonterminal(bnf_list)

	for _, bnf := range bnf_list {
		struct_element := []generator.StructElement{}

		right_set := map[string]bool{}
		for _, right := range bnf.Right {
			for _, right_ele := range right {
				right_set[right_ele] = true
			}
		}

		for right_ele, _ := range right_set {
			regex := regex.NewRegex("[a-zA-Z_]+")
			if match, ifmatch := regex.Match("0034"); ifmatch {

				if _, ok := nonterminals[right_ele]; ok {
					struct_element = append(struct_element, generator.StructElement{Name: right_ele, Type: "*" + right_ele})
				} else {
					struct_element = append(struct_element, generator.StructElement{Name: right_ele, Type: "*Terminal"})
				}
			}
		}

		Struct := generator.Struct{Name: bnf.Left, Elements: struct_element}
		structs = append(structs, Struct)
	}

	Struct := generator.Struct{Name: "Terminal", Elements: []generator.StructElement{generator.StructElement{Name: "Value", Type: "string"}}}
	structs = append(structs, Struct)

	return structs
}

func CreateNestedIf(bnf util.Bnf) []generator.Statement {
	nested_ifstatement := []generator.Statement{}
	for _, right := range bnf.Right {
		for _, right_ele := range right {
			body := generator.Statement{IfStatement: &generator.IfStatement{BoolExpression: &generator.BooleanExpression{Left: "right", Right: right_ele}}}
			nested_ifstatement = append(nested_ifstatement, body)
		}
	}

	return nested_ifstatement
}

func CreateFunction(bnf_list []util.Bnf) []generator.Function {
	functions := []generator.Function{}

	constructparser := generator.Function{}
	constructparser.Name = "ContructParserNode"

	var rootifstatement *generator.IfStatement
	var ifstatement *generator.IfStatement

	for i, bnf := range bnf_list {
		if i == 0 {
			nested_ifstatement := CreateNestedIf(bnf)
			rootifstatement = &generator.IfStatement{BoolExpression: &generator.BooleanExpression{Left: "root.String()", Right: bnf.Left}, Body: nested_ifstatement}
			ifstatement = rootifstatement
		} else {

			nested_ifstatement := CreateNestedIf(bnf)
			newifstatement := &generator.IfStatement{BoolExpression: &generator.BooleanExpression{Left: "root.String()", Right: bnf.Left}, Body: nested_ifstatement}
			ifstatement.Else = newifstatement
			ifstatement = newifstatement
		}
	}

	constructparser.Statements = []generator.Statement{generator.Statement{IfStatement: rootifstatement}}

	functions = append(functions, constructparser)

	return functions
}

func main() {

	_, filename, _, _ := runtime.Caller(0)
	bnf_path := filepath.Join(filepath.Dir(filepath.Dir(filepath.Dir(filename))), "sample2.bnf")

	bnf, err := ioutil.ReadFile(bnf_path)
	util.Check(err)

	bnf_parsed := util.Parse_bnf_file(string(bnf))
	fmt.Printf("%+v", bnf_parsed)

	structs := CreateStruct(bnf_parsed)
	functions := CreateFunction(bnf_parsed)

	package_temp := generator.Package{Name: "generator", Structs: structs, Functions: functions}

	f, err := os.Create("generated/sample.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	generator.Execute(f, package_temp)

}
