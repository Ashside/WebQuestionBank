package api

import (
	"fmt"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	// 在这里，你可以使用username和password来进行登录验证
	// 如果验证成功，你可以设置session或者返回一个token
	// 如果验证失败，你可以返回一个错误信息

	_, _ = fmt.Println("Username: ", username)
	_, _ = fmt.Println("Password: ", password)
}
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "处理器: %s\n", "registerHandler")
	if err != nil {
		return
	}

}
