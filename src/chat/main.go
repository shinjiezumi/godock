package chat

import (
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

type TemplateHandler struct {
	Once     sync.Once
	Filename string
	Templ    *template.Template
}

func (t *TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.Once.Do(func() {
		t.Templ = template.Must(template.ParseFiles(filepath.Join("./chat/templates", t.Filename)))
	})
	t.Templ.Execute(w, r)
}
