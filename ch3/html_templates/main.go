package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
)

// Article aaa
type Article struct {
	Name       string `json:"name,omitempty"`
	AuthorName string `json:"author_name,omitempty"`
	Draft      bool   `json:"draft,omitempty"`
}

// ArticleCollection aaa
type ArticleCollection struct {
	Articles []Article `json:"articles"`
	Total    int       `json:"total"`
}

// Product aaa
type Product struct {
	Price    float64
	Quantity float64
}

// Test custom JSON keys
type Test struct {
	// Field appears in JSON with the key "name".
	Name string `json:"name"`
	// Field appears in the JSON with the key "author_name",
	// but doesn’t appear at all if its value is empty.
	AuthorName string `json:"author_name,omitempty"`
	// Field will not appear in the JSON representation.
	CommissionPrice float64 `json:"-"`
}

// ByLine aaa
func (a Article) ByLine() string {
	return fmt.Sprintf("Written by %s", a.AuthorName)
}

// Multiply takes two float arguments and returns their multiplied value
func Multiply(a, b float64) float64 {
	return a * b
}

// Config aaa
type Config struct {
	Name     string `json:"SiteName"`
	URL      string `json:"SiteUrl"`
	Database struct {
		Name     string
		Host     string
		Port     int
		Username string
		Password string
	}
}

// FooJSON aaa
func FooJSON(input string) {
	data := map[string]interface{}{}
	err := json.Unmarshal([]byte(input), &data)
	if err != nil {
		panic(err)
	}

	foo, _ := data["foo"]

	switch foo.(type) {
	case float64:
		fmt.Printf("Float %f\n", foo)
	case string:
		fmt.Printf("String %s\n", foo)
	default:
		fmt.Printf("Something else\n")
	}
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
		Draft:      true,
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

	// Loops
	a := map[int]Article{
		1: goArticle,
		2: Article{Name: "blabla", AuthorName: "me"},
	}
	tmpl, err = template.New("Loops").Parse(`
    {{range .}}
        <p>{{.Name}} by {{.AuthorName}}</p>
    {{else}}
        <p>No published articles yet</p>
    {{end}}
    `)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, a)
	if err != nil {
		panic(err)
	}

	// Multiple templates
	tmpl, err = template.New("Multiple Templates").Parse(`
	{{define "ArticleResource"}}
	    <p>{{.Name}} by {{.AuthorName}}</p>
	{{end}}
	{{define "ArticleLoop"}}
	    {{range .}}
	        {{template "ArticleResource" .}}
	    {{else}}
	        <p>No published articles yet</p>
	    {{end}}
	{{end}}

    {{template "ArticleLoop" .}}
	`)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, a)
	if err != nil {
		panic(err)
	}

	// Pipelines
	tmpl, _ = template.New("Pipeline").Parse(
		"Price: ${{printf \"%.2f\" .}}\n",
	)
	tmpl.Execute(os.Stdout, 12.3)

	tmpl, _ = template.New("Pipeline").Parse(
		"Price: ${{. | printf \"%.2f\"}}\n",
	)
	tmpl.Execute(os.Stdout, 12.3)

	tmpl = template.New("Pipeline functions")
	tmpl.Funcs(template.FuncMap{"multiply": Multiply})
	tmpl, err = tmpl.Parse(`
    {{$total := multiply .Price .Quantity}}
    Price: ${{printf "%.2f" $total}}
    `)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, Product{12.30, 2})
	if err != nil {
		panic(err)
	}

	// JSON
	// Marshaling structs
	data, err := json.Marshal(goArticle)
	if err != nil {
		fmt.Println("Couldn't marshal article:", err)
	} else {
		fmt.Println(string(data))
	}

	data, _ = json.MarshalIndent(article, "", "  ")
	fmt.Println(string(data))

	// Nested types
	p1 := Article{Name: "JSON in Go"}
	p2 := Article{Name: "Marshaling is easy"}
	articles := []Article{p1, p2}
	collection := ArticleCollection{
		Articles: articles,
		Total:    len(articles),
	}
	data, err = json.MarshalIndent(collection, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	// Unmarshaling types
	conf := Config{}
	data, err = ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &conf)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Site: %s (%s)\n", conf.Name, conf.URL)

	db := conf.Database
	fmt.Printf(
		"DB: mysql://%s:%s@%s:%d/%s\n",
		db.Username,
		db.Password,
		db.Host,
		db.Port,
		db.Name,
	)

	// Unknown JSON
	FooJSON(`{
		"foo": 123
	}`)
	FooJSON(`{
		"foo": "bar"
	}`)
	FooJSON(`{
		"foo": []
	}`)
}
