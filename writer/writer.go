package writer

import (
	"fmt"
	"os"
	"text/template"

	"github.com/mrueg/go-deconstruct/types"
)

const modtmpl = `
{{- if .Module.Name -}}
module {{ .Module.Name }}

go {{ .GoRelease.Major }}.{{ .GoRelease.Minor }}
{{- end }}
{{- if .Dependencies }}

require (
{{- range $dependency := .Dependencies }}
	{{ $dependency.Name }} {{ $dependency.Version }}
{{- end}}
)
{{- end }}
{{- if .Replacements }}

replace (
{{- range $replacement := .Replacements }}
	{{ $replacement.Name }} => {{ $replacement.ReplacedWith }} {{ $replacement.Version }}
{{- end}}
)
{{- end }}
`

func WriteMod(modFile types.ModFile, outputPath string) error {
	var outputFile = os.Stdout
	tmpl, err := template.New("").Parse(modtmpl)
	if err != nil {
		panic(err)
	}
	if outputPath != "" {
		outputFile, err = os.Create(outputPath)
		if err != nil {
			return fmt.Errorf("Unable to create file: %s", err)
		}
	}
	err = tmpl.Execute(outputFile, modFile)
	if err != nil {
		panic(err)
	}
	return nil
}
