package server

//
// //TODO: 重写server包
//
// import (
// 	"github.com/xmh1011/IssueReport/github"
// 	"html/template"
// 	"log"
// 	"net/http"
// )
//
// func WebServer() { // 用来启动 web 服务
// 	http.HandleFunc("/", Handle)             // 设置访问的路由
// 	err := http.ListenAndServe(":8000", nil) // 设置监听的端口
// 	if err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// }
//
// // Handle is the handler for the web server
// func Handle(w http.ResponseWriter, r *http.Request) {
// 	// 定义一个模板，用来展示结果
// 	// template.Must 是一个辅助函数，用来检查模板是否有错误，如果有错误，会抛出异常。
// 	// template.New 是一个辅助函数，用来创建一个模板。
// 	var issueList = template.Must(template.New("issuelist").Parse(`
// <h1>{{.TotalCount}} issues</h1>
// <table>
// <tr style='text-align: left'>
//   <th>#</th>
//   <th>State</th>
//   <th>User</th>
//   <th>Title</th>
// </tr>
// {{range .Items}}
// <tr>
//   <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
//   <td>{{.State}}</td>
//   <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
//   <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
// </tr>
// {{end}}
// </table>
// `))
// 	issueList.Execute(w, result) // 执行模板
// }
