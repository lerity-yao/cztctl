package rabbitmqGen

import (
	"cztctl/config"
	"cztctl/pkg/spec"
	"cztctl/util/format"
	"cztctl/util/pathx"
	_ "embed"
)

const contextFilename = "service_context"

//go:embed svc.tpl
var contextTemplate string

func genServiceContext(dir, rootPkg string, cfg *config.Config, api *spec.ApiSpec) error {
	filename, err := format.FileNamingFormat(cfg.NamingFormat, contextFilename)
	if err != nil {
		return err
	}
	
	configImport := "\"" + pathx.JoinPackages(rootPkg, configDir) + "\""

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          contextDir,
		filename:        filename + ".go",
		templateName:    "contextTemplate",
		category:        category,
		templateFile:    contextTemplateFile,
		builtinTemplate: contextTemplate,
		data: map[string]string{
			"configImport": configImport,
			"config":       "config.Config",
		},
	})
}
