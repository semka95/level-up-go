package main

import (
	"fmt"
	"html/template"
	"os"
)

// Article aaa
type Article struct {
	Name       string
	AuthorName string
	Draft      bool
}

// ByLine aaa
func (a Article) ByLine() string {
	return fmt.Sprintf("Written by %s", a.AuthorName)
}

func main() {
	tmpl, err := template.New("Foo").Parse("<h1>Hello {{.}}</h1>\n")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, "World")
	if err != nil {
		panic(err)
	}

	// Accessing data
	goArticle := Article{
		Name:       "The Go html/template package",
		AuthorName: "Mal Curtis",
		//Draft:      true,
	}
	tmpl, err = template.New("Article Struct").Parse("'{{.Name}}' by {{.AuthorName}}\n")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, goArticle)
	if err != nil {
		panic(err)
	}

	article := map[string]string{
		"Name":       "The Go html/template package",
		"AuthorName": "Mal Curtis",
	}
	tmpl, err = template.New("Article Map").Parse("'{{.Name}}' by {{.AuthorName}}\n")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, article)
	if err != nil {
		panic(err)
	}

	tmpl, err = template.New("Article Function").Parse("{{.ByLine}}\n")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, goArticle)
	if err != nil {
		panic(err)
	}

	// Conditionals
	tmpl, err = template.New("Cond").Parse("{{.Name}}{{if .Draft}} (Draft){{else}} (Published){{end}}\n")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, goArticle)
}
