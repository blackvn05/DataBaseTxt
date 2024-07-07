package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func serveFile(w http.ResponseWriter, r *http.Request) {
	filePath := "Qs.txt" // Đường dẫn tới file Qs.txt
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	// Chuyển đổi nội dung file thành các dòng
	lines := strings.Split(string(content), "\n")

	// Hiển thị các dòng thô (raw lines)
	fmt.Fprintf(w, "<html><body><pre>")
	for _, line := range lines {
		fmt.Fprintf(w, "%s\n", line)
	}
	fmt.Fprintf(w, "</pre></body></html>")
}

func main() {
	http.HandleFunc("/Qs.txt", serveFile) // Định nghĩa route để phục vụ file Qs.txt

	fmt.Println("Server is running on http://localhost:0501")
	http.ListenAndServe(":0501", nil) // Chạy server trên cổng 0501
}
