package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Runner CommandLineInterface = &command{}
)

type CommandLineInterface interface {
	RootCmd() *cobra.Command
	GenerateCmd(cli *cobra.Command, args []string)
	DeployCmd(cli *cobra.Command, args []string)
}

type command struct{}

var rootCmd = cobra.Command{
	Use:  "root",
	Long: "this will run the entire application",
	Run: func(cli *cobra.Command, args []string) {
		fmt.Println("start application")
	},
}

func (command) RootCmd() *cobra.Command {
	rootCmd.AddCommand(GenerateCmd)
	rootCmd.AddCommand(DeployCmd)
	return &rootCmd
}
