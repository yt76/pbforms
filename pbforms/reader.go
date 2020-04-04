package pbforms

import (
	"net/http"
	"reflect"
	"strconv"
)

type FormReader struct {
}

func NewFormReader() *FormReader {
	return new(FormReader)
}

func (r *FormReader) Parse(req *http.Request, f interface{}) {
	req.ParseForm()
	v := reflect.Indirect(reflect.ValueOf(f))
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		ft := t.Field(i)
		if ft.Type.Kind() != reflect.Ptr {
			continue
		}
		s := req.FormValue(ft.Name)
		if ft.Type.Elem().Kind() == reflect.String {
			fv := v.FieldByName(ft.Name)
			fv.Set(reflect.ValueOf(&s))
		}
		if ft.Type.Elem().Kind() == reflect.Bool {
			fv := v.FieldByName(ft.Name)
			var b *bool = nil
			if s != "" {
				tr := true
				b = &tr
			}
			fv.Set(reflect.ValueOf(b))
		}
		if ft.Type.Elem().Kind() == reflect.Int32 {
			fv := v.FieldByName(ft.Name)
			i, err := strconv.Atoi(s)
			if err == nil {
				var i32 int32 = int32(i)
				fv.Set(reflect.ValueOf(&i32))
			}
		}
	}
}
