package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xmh1011/IssueReport/cmd"
)

func main() {

	// Get the root command
	mainCmd := GetCommand()

	// Add the subcommands
	Cmd := cmd.GetCommand()
	mainCmd.AddCommand(Cmd)

	if err := mainCmd.Execute(); err != nil {
		fmt.Printf("Error : %+v\n", err)
	}

}

func GetCommand() *cobra.Command {
	c := &cobra.Command{
		Use:  "github",
		Long: "search the github issue",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd:   true,  // Disable the default command
			DisableNoDescFlag:   true,  // Disable the flag
			DisableDescriptions: true}, // Disable the description
	}

	return c
}
