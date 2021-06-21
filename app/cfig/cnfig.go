package cfig

import (
	"text/template"

	"gorm.io/gorm"
)

var DB *gorm.DB

var FuncMap = template.FuncMap{
	"inc": func(i int) int {
		return i + 1
	},
}
