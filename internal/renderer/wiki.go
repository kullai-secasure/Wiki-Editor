package renderer

import (
	"bytes"
	"text/template"

	"wikiflex/internal/models"
)

var pageTemplate = template.Must(template.New("wiki").Parse(`
<!DOCTYPE html>
<html>
<head><title>{{.Title}} - WikiFlex</title></head>
<body>
	<article>
		<h1>{{.Title}}</h1>
		<div class="content">{{.Content}}</div>
		<footer>Last edited by {{.Author}} on {{.UpdatedAt}}</footer>
	</article>
</body>
</html>
`))

func RenderPage(page *models.WikiPage) (string, error) {
	var buf bytes.Buffer
	err := pageTemplate.Execute(&buf, page)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
