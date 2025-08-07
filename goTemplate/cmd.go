package goTemplate

import (
	"cztctl/config"
	"cztctl/goTemplate/goGen"
	"cztctl/goTemplate/rabbitmqGen"
	"cztctl/internal/cobrax"
)

var (
	// Cmd describes an api command.
	Cmd         = cobrax.NewCommand("go", cobrax.WithRunE(goGen.CreateGoTemplate))
	rabbitmqCmd = cobrax.NewCommand("rabbitmq", cobrax.WithRunE(rabbitmqGen.RabbitmqCommand))
)

func init() {
	var (
		cmdFlags         = Cmd.Flags()
		rabbitmqCmdFlags = rabbitmqCmd.Flags()
	)

	cmdFlags.StringVar(&goGen.VarStringHome, "home")

	rabbitmqCmdFlags.StringVar(&rabbitmqGen.VarStringDir, "dir")
	rabbitmqCmdFlags.StringVar(&rabbitmqGen.VarStringRabbitmq, "rabbitmq")
	rabbitmqCmdFlags.StringVar(&rabbitmqGen.VarStringHome, "home")
	rabbitmqCmdFlags.StringVar(&rabbitmqGen.VarStringRemote, "remote")
	rabbitmqCmdFlags.StringVar(&rabbitmqGen.VarStringBranch, "branch")
	rabbitmqCmdFlags.BoolVar(&rabbitmqGen.VarBoolWithTest, "test")
	rabbitmqCmdFlags.BoolVar(&rabbitmqGen.VarBoolTypeGroup, "type-group")
	rabbitmqCmdFlags.StringVarWithDefaultValue(&rabbitmqGen.VarStringStyle, "style", config.DefaultFormat)

	// Add sub-commands
	Cmd.AddCommand(rabbitmqCmd)
}
