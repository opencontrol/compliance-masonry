## Simple Google Go (golang) library for building templates for generic content
[![Go Report Card](https://goreportcard.com/badge/github.com/opencontrol/doc-template)](https://goreportcard.com/report/github.com/opencontrol/doc-template)

```go
func main() {
	funcMap := template.FuncMap{"title": strings.Title}
	docTemp, _ := GetTemplate("docx/fixtures/test.docx")
	docTemp.AddFunctions(funcMap)
	docTemp.Parse()
	docTemp.Execute("test.docx", nil)
}
```
