package main

import (
	"bytes"
	"fmt"
	"../forms"
	"github.com/golang/protobuf/proto"
	"html"
	"net/http"
	"../pbforms"
	"strings"
)

type Server struct {
	writer *pbforms.FormWriter
	reader *pbforms.FormReader
}

func NewServer() {
	sv := Server{}
	form := forms.MyForm{}
	sv.writer = pbforms.NewFormWriter(form)
	sv.reader = pbforms.NewFormReader()
	http.HandleFunc("/", sv.handle)
}

func (sv *Server) handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	if r.Method == "GET" {
		fmt.Fprintf(w, "<html><body>\n")
		sv.writer.Write("/", w)
		fmt.Fprintf(w, "</body></html>\n")
	}
	if r.Method == "POST" {
		fmt.Fprintf(w, "<html><body>\n")
		f := forms.MyForm{}
		sv.reader.Parse(r, &f)
		buf := new(bytes.Buffer)
		proto.MarshalText(buf, &f)
		s := html.EscapeString(buf.String())
		s = strings.Replace(s, "\n", "<br/>\n", -1)
		fmt.Fprintf(w, s + "<br/>\n")
		fmt.Fprintf(w, "<a href=\"/\">Back</a><br/>\n")
		fmt.Fprintf(w, "</body></html>\n")
	}
}

func main() {
	NewServer()
	fmt.Print("Serving at 8080\n")
	http.ListenAndServe(":8080", nil)
	fmt.Print("Stop serving\n")
}
