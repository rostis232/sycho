package service

import (
	// "errors"
	// "fmt"
	"html/template"
	"io"
	// "path/filepath"

	"github.com/labstack/echo/v4"
)

// type Template struct {
//     templates map[string]*template.Template
// }

// var functions = template.FuncMap{}

// func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
// 	tmpl, ok := t.templates[name]
// 	if !ok {
// 	  err := errors.New("Template not found -> " + name)
// 	  fmt.Println(err)
// 	  return err
// 	}
// 	return tmpl.ExecuteTemplate(w, name, data)
// }


// // RenderTemplate renders templates using html/template
// func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
// 	var tc map[string]*template.Template

// 	t, ok := tc[tmpl]
// 	if !ok {
// 		log.Fatal("Couldn`t get template in template cache with name: ", tmpl)
// 	}

// 	buf := new(bytes.Buffer)

// 	td = AddDefaultData(td, r)

// 	_ = t.Execute(buf, td)

// 	_, err := buf.WriteTo(w)
// 	if err != nil {
// 		fmt.Println("Error writing template to browser", err)
// 	}

// 	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
// 	err = parsedTemplate.Execute(w, nil)
// 	if err != nil {
// 		/*FIXME: Check the error. Template parsing is successful bot err != nil!
// 		(html/template:majors.page.tmpl:1:11: no such template)*/
// 		fmt.Println("error parsing template", err)
// 		return
// 	}
// }

// func NewTemplateSet () (*Template, error) {
// 	templates := make(map[string]*template.Template)
// 	tmpl, err := template.New("index").ParseFiles("./templates/index.page.html", "./templates/base.layout.html")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	templates["index"] = tmpl
// 	return &Template{
// 		templates: templates,
// 	}, nil
// }

// func NewTemplateSet () (*Template, error) {
// 	myCache := map[string]*template.Template{}
// 	pages, err := filepath.Glob("./templates/*.page.tmpl")
// 	// fmt.Println(pages)
// 	if err != nil {
// 		return &Template{}, err
// 	}

// 	for _, page := range pages {
// 		name := filepath.Base(page)
// 		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
// 		if err != nil {
// 			fmt.Println("Parsing files error", err)
// 		}
// 		// fmt.Println("Page is currently", page)

// 		matches, err := filepath.Glob("./templates/*.layout.html")
// 		if err != nil {
// 			return &Template{}, err
// 		}

// 		if len(matches) > 0 {
// 			ts, err = ts.ParseGlob("./templates/*.layout.html")
// 			if err != nil {
// 				return &Template{}, err
// 			}
// 		}

// 		myCache[name] = ts
// 	}

// 	return &Template{templates: myCache}, nil

// }

type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var Tmps = &Template{
    templates: template.Must(template.ParseGlob("templates/*.html")),
}