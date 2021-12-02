package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	html := "<html>"
	html += "<body>"
	html += "<h1>Hola Mundo</h1>"
	html += "</body>"
	html += "</html>"

	w.Write([]byte(html))
}

var productos []string

func producto(w http.ResponseWriter, r *http.Request) {

	/* r.ParseForm()
	add, okForm := r.Form["add"]
	if okForm && len(add) == 1 {
		productos = append(productos, string(add[0]))
		w.Write([]byte("Producto añadido correctamente"))
		return
	} */

	html := "<html>"
	html += "<body>"
	html += "<h1>Total de Productos: " + strconv.Itoa(len(productos)) + "</h1>"
	html += "</body>"
	html += "</html>"

	w.Write([]byte(html))
}

func main() {

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", home)
	http.HandleFunc("/producto", producto)
	http.HandleFunc("/info", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "Host: ", req.Host)
		fmt.Fprintln(w, "URI: ", req.RequestURI)
		fmt.Fprintln(w, "Method: ", req.Method)
		fmt.Fprintln(w, "RemoteAddr: ", req.RemoteAddr)

	})
	http.ListenAndServe(":8080", nil)
}
