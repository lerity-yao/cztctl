package rabbitmqGen

import (
	apiParser "github.com/lerity-yao/cztctl/pkg/parser"
	"github.com/lerity-yao/cztctl/pkg/parser/ast"
	"github.com/lerity-yao/cztctl/pkg/spec"
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
