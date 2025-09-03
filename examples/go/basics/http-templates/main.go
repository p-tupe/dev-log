// This code showcases a project structure using go templates
// https://pkg.go.dev/html/template#example-Template-Parsefiles
package main

import (
	"html/template"
	"log/slog"
	"net/http"
	"os"
)

type PageData struct {
	Title string
	Page  string
}

func main() {
	tmplMap := genAppTmpl()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[1:]
		if path == "" {
			path = "index"
		}
		path += ".tmpl"

		// We can pick a template to execute by its name
		err := tmplMap[path].ExecuteTemplate(w, path, nil)
		if err != nil {
			slog.Error(err.Error())
		}
	})

	http.ListenAndServe(":8080", nil)

	// This code won't run as is,
	// just move it up to see what it does.
	simpleTmpl()
}

func genAppTmpl() map[string]*template.Template {
	// We don't need a name here, instead each file name is taken as a template name
	// tmpl := template.Must(template.ParseGlob("./templates/**/*.tmpl"))
	// slog.Info(tmpl.DefinedTemplates())

	// However, if we use a root layout, then we must parse all templates
	// separately to differentiate the "main" content of the file.
	// We just loop over each page and parse it using the layout dir.
	tmplMap := make(map[string]*template.Template)

	files, err := os.ReadDir("./templates/pages")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		name := file.Name()
		tmplMap[name] = template.Must(template.ParseFiles("./templates/pages/" + name))
		tmplMap[name] = template.Must(tmplMap[name].ParseGlob("./templates/components/*.tmpl"))
		tmplMap[name] = template.Must(tmplMap[name].ParseGlob("./templates/layouts/*.tmpl"))
	}

	return tmplMap
}

// This is the very basic of how a template works.
// It has some name, and expects some data in.
// template.Must is optional "panic on error" wrapper
func simpleTmpl() {
	tx := template.Must(template.New("some.tmpl").Parse("This is {{ . }} sent to the template!"))
	tx.Execute(os.Stdout, "some data")

}
