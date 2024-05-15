package api

import (
	"fmt"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "处理器: %s\n", "loginHandler")
	if err != nil {
		return
	}

}
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "处理器: %s\n", "registerHandler")
	if err != nil {
		return
	}

}
