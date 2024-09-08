// Copyright 2024 Michael F. Collins, III
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.

package template

import (
	"embed"
	"html/template"
	"io"
	"path"

	"github.com/mfcollins3/generate-github-readme/internal/template/functions/skills"
)

//go:embed templates
var templates embed.FS

type Generator struct {
	t    *template.Template
	name string
}

func NewGenerator(templatePath string) (*Generator, error) {
	templateName := path.Base(templatePath)
	t, err := template.New(templateName).
		Funcs(template.FuncMap{
			"ReadSkills": skills.ReadSkills,
		}).
		ParseFS(templates, "templates/*")
	if err != nil {
		return nil, err
	}

	t, err = t.ParseFiles(templatePath)
	if err != nil {
		return nil, err
	}

	return &Generator{
		t:    t,
		name: templateName,
	}, nil
}

func (g *Generator) Generate(w io.Writer) error {
	return g.t.ExecuteTemplate(w, g.name, "Test")
}
