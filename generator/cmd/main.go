package main

import (
	"fmt"
	"github.com/KazumaTakata/go_code_generator"
	"github.com/KazumaTakata/lr_parser/util"
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
			if _, ok := nonterminals[right_ele]; ok {
				struct_element = append(struct_element, generator.StructElement{Name: right_ele, Type: "*" + right_ele})
			} else {
				struct_element = append(struct_element, generator.StructElement{Name: right_ele, Type: "*Terminal"})
			}
		}

		Struct := generator.Struct{Name: bnf.Left, Elements: struct_element}
		structs = append(structs, Struct)
	}

	Struct := generator.Struct{Name: "Terminal", Elements: []generator.StructElement{generator.StructElement{Name: "Value", Type: "string"}}}
	structs = append(structs, Struct)

	return structs
}

func CreateFunction(bnf_list []util.Bnf) []generator.Function {
	functions := []generator.Function{}

	constructparser := generator.Function{}
	constructparser.Name = "ContructParserNode"

	var rootifstatement *generator.IfStatement
	var ifstatement *generator.IfStatement

	for i, bnf := range bnf_list {
		if i == 0 {
			rootifstatement = &generator.IfStatement{BoolExpression: &generator.BooleanExpression{Left: "root.String()", Right: bnf.Left}}
			ifstatement = rootifstatement
		} else {
			newifstatement := &generator.IfStatement{BoolExpression: &generator.BooleanExpression{Left: "root.String()", Right: bnf.Left}}

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

	generator.Execute(os.Stdout, package_temp)

}
