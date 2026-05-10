package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

type MyHandler struct {
	NextHandler http.Handler
}

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Max-Age", "86400")
	w.WriteHeader(http.StatusOK)
}

// 实现 http.Handler 接口的 ServeHTTP 方法
func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		enableCORS(w)
		return
	}

	reqPath := r.URL.Path
	if reqPath == "/files/upload" {
		//uploadHandler(w, r)
		return
	}

	if strings.HasPrefix(reqPath, "/files/static/") {
		getHandler(w, r)
		return
	}

	h.NextHandler.ServeHTTP(w, r)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request: ", r.URL.Path)
	if strings.HasPrefix(r.URL.Path, "/files/static/stream") {
		filePath := r.URL.Query().Get("path")
		// base64data := r.URL.Query().Get("path")

		// fmt.Println("Base64: ", base64data)
		// decodedBytes, err := base64.StdEncoding.DecodeString(base64data)

		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusBadRequest)
		// 	return
		// }

		// filePath := string(decodedBytes)
		fmt.Println("Filepath: ", filePath)

		if filePath == "" {
			http.Error(w, "no path", http.StatusBadRequest)
			return
		}

		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}

		fmt.Println("Serve the file")
		http.ServeFile(w, r, filePath)
		fmt.Println("Exit")
		return
	}
	http.NotFound(w, r)
}
