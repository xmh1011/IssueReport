package main

import (
	"github.com/spf13/cobra"
	"github.com/xmh1011/IssueReport/cmd"
	"github.com/xmh1011/IssueReport/server"
	"net/http"
)

func main() {
	
	mainCmd := GetCommand()
	
	Cmd := cmd.GetCommand()
	mainCmd.AddCommand(Cmd)
	
	http.HandleFunc("/", server.Handle)      // 设置访问的路由
	http.ListenAndServe("0.0.0.0:8080", nil) // 设置监听的端口
	
}

func GetCommand() *cobra.Command {
	c := &cobra.Command{
		Use:  "github",
		Long: "search the github issue",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd:   true,
			DisableNoDescFlag:   true,
			DisableDescriptions: true},
	}
	
	return c
}
