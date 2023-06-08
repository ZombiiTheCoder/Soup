package spoon

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func init_pkg(name, desc string) {
	rgx:=regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	id:=rgx.ReplaceAllString(strings.TrimSpace(name), "")
	m, _:=os.Create("main.soup")
	m.WriteString(`
	
	func main() {

		println("Hello World")

	}

	main()

	`)
	f, _:=os.Create("pkg.json")
	f.WriteString(fmt.Sprintf(`
	{

		"name": "%v",
		"id": "%v",
		"description": "%v",
		"dependencies": []
	
	}
	`, name, id, desc))
}