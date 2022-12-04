package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xmh1011/IssueReport/github"
	"github.com/xmh1011/IssueReport/server"
	"io"
	"log"
	"os"
)

type Options struct {
	Stderr io.Writer // Stderr is the writer to write warnings and errors to.
	Stdout io.Writer // Stdout is the writer to write normal output to.
	
	repo   string // Specify the repository
	status string // Specify the status of issue
	key    string // Specify the key of issue
	web    bool   // Specify whether to open the web server
}

// NewOptions returns an Options struct with default values set.
func NewOptions() *Options {
	return &Options{
		Stderr: os.Stderr,
		Stdout: os.Stdout,
	}
}

var opt = NewOptions()

// GetCommand returns the cobra command for issue
func GetCommand() *cobra.Command {
	// Main command
	c := &cobra.Command{ // c is the main command
		Use:   "issue", // The name of the command
		Short: "Create a new issue on GitHub",
		Run: func(cmd *cobra.Command, args []string) {
			commands := make([]string, 0)
			if opt.repo != "" { // If the repo is not empty, add the repo to the commands
				commands = append(commands, opt.repo)
			} else { // If the repo is empty, print the error message
				fmt.Printf("The repo name should be specified!\n")
			}
			if opt.status != "" {
				commands = append(commands, opt.status)
			}
			if opt.key != "" {
				commands = append(commands, opt.key)
			}
			result, err := github.SearchIssues(commands) // 调用SearchIssues函数
			if err != nil {
				log.Fatal(err)
			}
			if opt.web {
				server.WebServer()
			}
			github.IssueReport(result, err)
		},
	}
	// Add flags
	c.Flags().StringVarP(&opt.repo, "repo", "r", "", "Specify the repository") // Add the flag of repo
	c.Flags().StringVarP(&opt.status, "status", "s", "is:open", "Specify the status")
	c.Flags().StringVarP(&opt.key, "key", "k", "", "Specify the key")
	return c
}
