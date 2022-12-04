package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const IssueUrl = "https://api.github.com/search/issues" // IssueUrl是GitHub提供的issue跟踪接口

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues函数查询GitHub的issue跟踪接口
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	// QueryEscape转义字符串，以便将其安全地放在URL查询中
	q := url.QueryEscape(strings.Join(terms, " ")) // 将terms中的元素用空格连接起来，并进行url编码
	resp, err := http.Get(IssueUrl + "?q=" + q)    // 发送get请求
	if err != nil {
		return nil, err
	}
	
	// We must close resp.Body on all execution paths.
	if resp.StatusCode != http.StatusOK { // 如果请求失败，则返回错误信息
		resp.Body.Close()                                              // 关闭请求
		return nil, fmt.Errorf("search query failed: %s", resp.Status) // fmt.Errorf 会返回一个error
	}
	
	// The json package can decode an io.Reader value that contains JSON data into a variable.
	var result IssuesSearchResult                                      // 定义一个IssuesSearchResult类型的变量
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil { // 将resp.Body中的json数据解析到result中
		resp.Body.Close() // 关闭请求
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
