package main

import (
	"github.com/spf13/cobra"
	"github.com/xmh1011/IssueReport/cmd"
)

func main() {
	
	// Get the root command
	mainCmd := GetCommand()
	
	// Add the subcommands
	Cmd := cmd.GetCommand()
	mainCmd.AddCommand(Cmd)
	
}

func GetCommand() *cobra.Command {
	c := &cobra.Command{
		Use:  "github",
		Long: "search the github issue",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd:   true, // Disable the default command
			DisableNoDescFlag:   true, // Disable the flag
			DisableDescriptions: true}, // Disable the description
	}
	
	return c
}
