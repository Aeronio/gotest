package config

import "html/template"

//Notice this capital letters, this means we'll be using it in different packages
var TPL *template.Template

func init() {
	TPL = template.Must(template.ParseGlob("templates/*.gohtml"))
}
