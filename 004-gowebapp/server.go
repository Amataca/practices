package main

import (
	"fmt"
	"net/http"
)

func main(){

    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    fmt.Println(&fs)
    http.ListenAndServe(":8080", nil)
}
