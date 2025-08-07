package goGen

import (
	"fmt"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var (
	// VarStringHome describes the output.
	VarStringHome string
)

// CreateGoTemplate create api template file
func CreateGoTemplate(_ *cobra.Command, _ []string) error {
	fmt.Println(color.Green.Render("Done."))
	return nil
}
