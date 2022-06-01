package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var DeployCmd = &cobra.Command{
	Use:  "deploy",
	Long: "use for deploying smart contracts at parking addresses",
	Run:  Runner.DeployCmd,
}

func (command) DeployCmd(cli *cobra.Command, args []string) {
	fmt.Println("Deploying")
}
