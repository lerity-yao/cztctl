package analyzer

import (
	"cztctl/goTemplate/rabbitmqGen/parser"
	"fmt"
)

// Parse parses the given file and returns the parsed spec.
func Parse(filename string, src interface{}) {
	p := parser.New(filename, src)
	fmt.Println(p)
	ast := p.Parse()
	fmt.Println(ast)
	//if err := p.CheckErrors(); err != nil {
	//	return nil, err
	//}
	//
	//is := importstack.New()
	//err := is.Push(ast.Filename)
	//if err != nil {
	//	return nil, err
	//}
	//
	//importSet := map[string]lang.PlaceholderType{}
	//api, err := convert2API(ast, importSet, is)
	//if err != nil {
	//	return nil, err
	//}
	//if err := api.SelfCheck(); err != nil {
	//	return nil, err
	//}
	//
	//var result = new(spec.ApiSpec)
	//analyzer := Analyzer{
	//	api:  api,
	//	spec: result,
	//}
	//
	//err = analyzer.convert2Spec()
	//if err != nil {
	//	return nil, err
	//}

}
