package rabbitmqGen

import (
	"errors"
	"fmt"
	"github.com/gookit/color"
	"github.com/lerity-yao/cztctl/config"
	"github.com/lerity-yao/cztctl/pkg/golang"
	"github.com/lerity-yao/cztctl/util"
	"github.com/lerity-yao/cztctl/util/pathx"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
	"path"
)

var (
	// VarStringDir describes a directory.
	tmpDir = path.Join(os.TempDir(), "goctl")
	// VarStringDir describes the directory.
	VarStringDir string
	// VarStringRabbitmq describes the API.
	VarStringRabbitmq string
	// VarStringHome describes the go home.
	VarStringHome string
	// VarStringRemote describes the remote git repository.
	VarStringRemote string
	// VarStringBranch describes the branch.
	VarStringBranch string
	// VarStringStyle describes the style of output files.
	VarStringStyle  string
	VarBoolWithTest bool
	// VarBoolTypeGroup describes whether to group types.
	VarBoolTypeGroup bool
)

// RabbitmqCommand gen rabbitmq go project files from command line
func RabbitmqCommand(_ *cobra.Command, _ []string) error {
	rabbitmqFile := VarStringRabbitmq
	dir := VarStringDir
	namingStyle := VarStringStyle
	home := VarStringHome
	remote := VarStringRemote
	branch := VarStringBranch
	withTest := VarBoolWithTest
	if len(remote) > 0 {
		repo, _ := util.CloneIntoGitHome(remote, branch)
		if len(repo) > 0 {
			home = repo
		}
	}

	if len(home) > 0 {
		pathx.RegisterCztctlHome(home)
	}
	if len(rabbitmqFile) == 0 {
		return errors.New("missing -rabbitmq")
	}
	if len(dir) == 0 {
		return errors.New("missing -dir")
	}

	return DoGenProject(rabbitmqFile, dir, namingStyle, withTest)
}

// DoGenProject gen go project files with rabbitmq file
func DoGenProject(rabbitmqFile, dir, style string, withTest bool) error {
	api, err := Parse(rabbitmqFile)
	if err != nil {
		return err
	}

	if err := api.Validate(); err != nil {
		return err
	}

	cfg, err := config.NewConfig(style)
	if err != nil {
		return err
	}
	logx.Must(pathx.MkdirIfNotExist(dir))
	rootPkg, err := golang.GetParentPackage(dir)
	if err != nil {
		return err
	}
	logx.Must(genEtc(dir, cfg, api))
	logx.Must(genConfig(dir, cfg, api))
	logx.Must(genMain(dir, rootPkg, cfg, api))
	logx.Must(genServiceContext(dir, rootPkg, cfg, api))
	logx.Must(genTypes(dir, cfg, api))
	logx.Must(genListener(dir, rootPkg, cfg, api))
	logx.Must(genHandlers(dir, rootPkg, cfg, api))
	logx.Must(genLogic(dir, rootPkg, cfg, api))

	fmt.Println(color.Green.Render("Done."))
	return nil
}
