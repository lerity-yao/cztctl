package rabbitmqGen

import (
	"cztctl/pkg/spec"
	"cztctl/util/format"
	"cztctl/vars"
	_ "embed"
	"fmt"
	"github.com/lerity-yao/cztctl/config"
	"strings"
)

const (
	configFile = "config"

	jwtTemplate = ` struct {
		AccessSecret string
		AccessExpire int64
	}
`
	jwtTransTemplate = ` struct {
		Secret     string
		PrevSecret string
	}
`
)

//go:embed config.tpl
var configTemplate string

func genConfig(dir string, cfg *config.Config, api *spec.ApiSpec) error {
	filename, err := format.FileNamingFormat(cfg.NamingFormat, configFile)
	if err != nil {
		return err
	}

	listenerAuthNames := getListenerAuths(api)

	//jwtTransNames := getJwtTrans(api)
	//var jwtTransList []string
	//for _, item := range jwtTransNames {
	//	jwtTransList = append(jwtTransList, fmt.Sprintf("%s %s", item, jwtTransTemplate))
	//}
	authImportStr := fmt.Sprintf("\"%s/rest\"", vars.ProjectOpenSourceURL)
	authImportStr = authImportStr + fmt.Sprintf("\n\"%s/go-mq/rabbitmq\"", vars.LerityOpenSourceUrl)

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          configDir,
		filename:        filename + ".go",
		templateName:    "configTemplate",
		category:        category,
		templateFile:    configTemplateFile,
		builtinTemplate: configTemplate,
		data: map[string]string{
			"authImport":         authImportStr,
			"listenerAuthConfig": strings.Join(listenerAuthNames, "\n"),
		},
	})
}
