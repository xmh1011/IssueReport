package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xmh1011/IssueReport/github"
	"io"
	"log"
	"os"
)

type Options struct {
	Stderr io.Writer
	Stdout io.Writer
	
	repo   string
	status string
	key    string
}

func NewOptions() *Options {
	return &Options{
		Stderr: os.Stderr,
		Stdout: os.Stdout,
	}
}

var opt = NewOptions()

func GetCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "issue",
		Short: "Create a new issue on GitHub",
		Run: func(cmd *cobra.Command, args []string) {
			commands := make([]string, 0)
			if opt.repo != "" {
				commands = append(commands, opt.repo)
			} else {
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
			fmt.Printf("%d issues:\n", result.TotalCount)
			for _, item := range result.Items {
				fmt.Printf("#%-5d %9.9s %s\n", item.Number, item.User.Login, item.Title)
			}
		},
	}
	c.Flags().StringVarP(&opt.repo, "repo", "r", "", "Specify the repository")
	c.Flags().StringVarP(&opt.status, "status", "s", "is:open", "Specify the status")
	c.Flags().StringVarP(&opt.key, "key", "k", "", "Specify the key")
	return c
}
