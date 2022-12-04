package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xmh1011/IssueReport/github"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// TODO: 改用接口
type Server interface {
	WebServer()
}

var Commands []string

type Options struct {
	Stderr io.Writer // Stderr is the writer to write warnings and errors to.
	Stdout io.Writer // Stdout is the writer to write normal output to.

	repo   string // Specify the repository
	status string // Specify the status of issue
	key    string // Specify the key of issue
	web    bool   // Specify whether to open the web server
	port   string // Specify the port of web server
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
		Use:   "issue",                   // The name of the command
		Short: "Search the github issue", // The short description of the command
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
			if opt.web { // If the web is true, open the web server
				Commands = commands
				opt.WebServer()
			} else { // If the web is false, print the result to the stdout
				github.IssueReport(result, err)
			}
		},
	}
	// Add flags
	c.Flags().StringVarP(&opt.repo, "repo", "r", "", "Specify the repository") // Add the flag of repo
	c.Flags().StringVarP(&opt.status, "status", "s", "is:open", "Specify the status")
	c.Flags().StringVarP(&opt.key, "key", "k", "", "Specify the key")
	c.Flags().BoolVarP(&opt.web, "web", "w", false, "Whether to open a web server (default: false)")
	c.Flags().StringVarP(&opt.port, "port", "p", "8000", "Specify the port of web server")
	return c
}

func (opt *Options) WebServer() { // 用来启动 web 服务
	http.HandleFunc("/", Handle) // 设置访问的路由
	args := ":" + opt.port
	err := http.ListenAndServe(args, nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// Handle is the handler for the web server
func Handle(w http.ResponseWriter, r *http.Request) {
	result, err := github.SearchIssues(Commands) // 调用 SearchIssues 函数
	// now 为现在的时间，yearAgo 为距现在一年的时间，monthAgo 为距现在一月的时间。
	now := time.Now()
	yearAgo := now.AddDate(-1, 0, 0)
	monthAgo := now.AddDate(0, -1, 0)

	// 三个切片，用来存储 不足一个月的问题，不足一年的问题，超过一年的问题。
	yearAgos := make([]*github.Issue, 0)
	monthAgos := make([]*github.Issue, 0)
	lessMonths := make([]*github.Issue, 0)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		// 如果 yearAgo 比 创建时间晚，说明超过一年
		if yearAgo.After(item.CreatedAt) {
			yearAgos = append(yearAgos, item)
			// 如果 monthAgo 比 创建时间晚，说明超过一月 不足一年
		} else if monthAgo.After(item.CreatedAt) {
			monthAgos = append(monthAgos, item)
			// 如果 monthAgo 比 创建时间早，说明不足一月。
		} else if monthAgo.Before(item.CreatedAt) {
			lessMonths = append(lessMonths, item)
		}
	}

	fmt.Fprintf(w, "\n一年前\n")
	for _, item := range yearAgos {
		fmt.Fprintf(w, "#%-5d %9.9s %.55s %v\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
	}

	fmt.Fprintf(w, "\n一月前\n")
	for _, item := range monthAgos {
		fmt.Fprintf(w, "#%-5d %9.9s %.55s %v\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt)
	}

	fmt.Fprintf(w, "\n不足一月\n")
	for _, item := range lessMonths {
		fmt.Fprintf(w, "#%-5d %9.9s %.55s %-40v\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
}
