package main

import (
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
	"strings"
	"text/template"
)

type MyHandler struct {
}

func getContentType(filename string) string {
	if strings.HasSuffix(filename, ".css") {
		return "text/css"
	} else if strings.HasSuffix(filename, ".html") {
		return "text/html"
	} else if strings.HasSuffix(filename, ".js") {
		return "application/javascript"
	} else if strings.HasSuffix(filename, ".mp4") {
		return "video/mp4"
	} else {
		return "text/plain"
	}
}

type IndexTemplateContext struct {
	FirstName string
	LastName string
	Message string
}


func (this *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var path = ""
	if r.URL.Path == "/" {
		path = "index.html"
	} else {
		path = r.URL.Path[1:]
	}
	log.Println(r.Method + ": " + path)
	data, err := ioutil.ReadFile(string(path))
	tmpl, err := template.New("anyName").Parse(string(data))

	if err == nil {
		contentType := getContentType(path)
		w.Header().Add("Content-Type", contentType)

		if contentType == "text/html" {
			context := IndexTemplateContext{"Cornelius", "Schmale", "Hello World"}
			tmpl.Execute(w, context)
		} else {
			w.Write(data)
		}
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404" + http.StatusText(404)))
	}

}

func main() {
	fmt.Println("running server at localhost:8080")

	http.Handle("/", new(MyHandler))
	http.ListenAndServe(":8080", nil)
}
