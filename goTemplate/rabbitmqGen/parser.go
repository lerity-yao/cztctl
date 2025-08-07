package rabbitmqGen

import (
	apiParser "cztctl/pkg/parser"
	"cztctl/pkg/spec"
	"github.com/zeromicro/go-zero/tools/goctl/api/parser/g4/ast"
)

type parser struct {
	ast  *ast.Api
	spec *spec.ApiSpec
}

// Parse parses the api file.
// Depreacted: use tools/goctl/pkg/parser/api/parser/parser.go:18 instead,
// it will be removed in the future.
func Parse(filename string) (*spec.ApiSpec, error) {
	return apiParser.Parse(filename, "")
}
