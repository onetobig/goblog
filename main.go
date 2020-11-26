package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Hello, 欢迎来到 goblog!</h1>")
	}
}
func main()  {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
