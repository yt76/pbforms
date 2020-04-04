package pbforms

import (
	"fmt"
	"io"
	"reflect"
)

type FormWriter struct {
	w io.Writer
	f interface{}
	path string
}

func NewFormWriter(f interface{}) *FormWriter {
	fw := new(FormWriter)
	fw.f = f
	return fw
}

func (fw *FormWriter) genField(name, valueType string) {
	t := "text"
	if valueType == "bool" {
		t = "checkbox"
	}
	fmt.Fprintf(fw.w, "%v<input type=\"%v\" name=\"%v\"/><br/>\n", name, t, name)
}

func (fw *FormWriter) WriteHeader(path string) {
	fmt.Fprintf(fw.w, "<form method=\"POST\" action=\"%s\">\n", path)
}

func (fw *FormWriter) WriteFooter() {
	fmt.Fprintf(fw.w, "<input type=\"submit\">\n")
	fmt.Fprintf(fw.w, "</form>\n")
}

func (fw *FormWriter) Write(path string, w io.Writer) {
	fw.w = w
	fw.WriteHeader(path)
	v := reflect.Indirect(reflect.ValueOf(fw.f))
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		ft := t.Field(i)
		if ft.Type.Kind() != reflect.Ptr {
			continue
		}
		if ft.Type.Elem().Kind() == reflect.Int32 {
			fw.genField(ft.Name, "int")
		}
		if ft.Type.Elem().Kind() == reflect.String {
			fw.genField(ft.Name, "string")
		}
		if ft.Type.Elem().Kind() == reflect.Bool {
			fw.genField(ft.Name, "bool")
		}
	}
	fw.WriteFooter()
}
