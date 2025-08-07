package rabbitmqGen

import (
	_ "embed"
	"fmt"
	"github.com/lerity-yao/cztctl/config"
	"github.com/lerity-yao/cztctl/pkg/spec"
	"github.com/lerity-yao/cztctl/util/format"
	"strconv"
	"strings"
)

const (
	defaultPort = 8888
	etcDir      = "etc"
)

//go:embed etc.tpl
var etcTemplate string

func genEtc(dir string, cfg *config.Config, api *spec.ApiSpec) error {
	filename, err := format.FileNamingFormat(cfg.NamingFormat, api.Service.Name)
	if err != nil {
		return err
	}

	service := api.Service
	host := "0.0.0.0"
	port := strconv.Itoa(defaultPort)

	listenerNames := getListenerConfig(api)
	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          etcDir,
		filename:        fmt.Sprintf("%s.yaml", filename),
		templateName:    "etcTemplate",
		category:        category,
		templateFile:    etcTemplateFile,
		builtinTemplate: etcTemplate,
		data: map[string]string{
			"serviceName":  service.Name,
			"host":         host,
			"port":         port,
			"listenConfig": strings.Join(listenerNames, "\n"),
		},
	})
}
