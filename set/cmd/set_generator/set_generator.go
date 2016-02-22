package main

import (
	"flag"
	"log"
	"os"
	"text/template"
	"unicode"
)

var setTemplate = `package {{.Pkg}}

import "github.com/jrozner/go-set/set"

type container struct {
	value interface{}
}

func (c *container) Value() interface{} {
	return c.value
}

func (c *container) Hash() interface{} {
	return ""
}

type {{.Name}} struct {
	s *set.{{.Implementation}}
}

func New{{.Name}}() *{{.Name}} {
	return &{{.Name}}{
		s: set.New{{.Implementation}}(),
	}
}

func (s *{{.Name}}) Add(item {{.rType}}) {
	s.s.Add(item)
}
`

func main() {
	var (
		implementation string
		name           string
		rtype          string
		pkg            string
	)

	flag.StringVar(&pkg, "package", "", "package name")
	flag.StringVar(&implementation, "implementation", "simple", "set implementation")
	flag.StringVar(&name, "name", "", "set name")
	flag.StringVar(&rtype, "type", "", "type")

	flag.Parse()

	if pkg == "" {
		log.Fatal("must specify a package for the set")
	}

	if name == "" {
		log.Fatal("must specify a name for the set")
	}

	if rtype == "" {
		log.Fatal("must specify a return type")
	}

	switch implementation {
	case "simple":
		implementation = "SimpleSet"
	case "orderd_simple":
		implementation = "OrderedSimpleSet"
	case "hash":
		implementation = "HashSet"
	case "ordered_hash":
		implementation = "OrderedHashSet"
	default:
		log.Fatal("must specify a valid implementation")
	}

	data := struct {
		Name           string
		Implementation string
		Pkg            string
	}{
		Name:           ToCamel(name, true),
		Implementation: implementation,
		Pkg:            pkg,
	}

	t := template.Must(template.New("set").Parse(setTemplate))

	fp, err := os.OpenFile(ToSnake(name)+"_set.go", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(fp, data)
}

func ToSnake(str string) string {
	chars := make([]rune, 0, len(str))

	for i, ch := range str {
		if unicode.IsUpper(ch) {
			ch = unicode.ToLower(ch)
			if i != 0 {
				chars = append(chars, '_')
			}
		}

		chars = append(chars, ch)
	}

	return string(chars)
}

func ToCamel(str string, export bool) string {
	capitalize := false
	chars := make([]rune, 0, len(str))

	for i, ch := range str {
		if ch == '_' {
			capitalize = true
			continue
		}

		if export == true && i == 0 && unicode.IsLower(ch) || capitalize == true {
			ch = unicode.ToUpper(ch)
		}

		chars = append(chars, ch)
		capitalize = false
	}

	return string(chars)
}
