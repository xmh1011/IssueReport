package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xmh1011/IssueReport/github"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
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
				WebServer()
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
	return c
}

func WebServer() { // 用来启动 web 服务
	http.HandleFunc("/", Handle)             // 设置访问的路由
	err := http.ListenAndServe(":8000", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// Handle is the handler for the web server
func Handle(w http.ResponseWriter, r *http.Request) {
	result, err := github.SearchIssues(Commands) // 调用 SearchIssues 函数
	if err != nil {
		log.Fatal(err) // 如果出错，打印错误信息
	}
	// 定义一个模板，用来展示结果
	// template.Must 是一个辅助函数，用来检查模板是否有错误，如果有错误，会抛出异常。
	// template.New 是一个辅助函数，用来创建一个模板。
	var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))
	issueList.Execute(w, result) // 执行模板
}
