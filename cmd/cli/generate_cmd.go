package cli

import (
	"github.com/spf13/cobra"
)

var GenerateCmd = &cobra.Command{
	Use:  "generate",
	Long: "use for geneating parking addresses",
	Run:  Runner.GenerateCmd,
}

func (command) GenerateCmd(cli *cobra.Command, args []string) {
	// fmt.Println("GenerateCmd")
	// input, _ := hex.DecodeString("0xff")
	// input = append(input, []byte("deplyingAddr")...)
	// input = append(input, []byte("salt")...)
	// input = append(input, core.HashKeccak256([]byte("bytecode"))...)
	// output := core.HashKeccak256(input)
	// fmt.Printf("%x\n", string(output))
	// fmt.Println(output)

}
