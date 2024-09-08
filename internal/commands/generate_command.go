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

package commands

import (
	readme_template "github.com/mfcollins3/generate-github-readme/internal/template"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var GenerateCommand = &cobra.Command{
	Use:     "generate-readme",
	Version: "0.1.0",
	Short:   "Generates a profile README for GitHub users",
	Long: `generate-readme implements a GitHub Action that is used to generate personal
profiles for GitHub users. The personal profile is created as a README.md
document that is created in a GitHub repository with the same name as the
user (ex. mfcollins3/mfcollins3). generate-readme will use a template to
dynamically generate the personal profile and will include dynamic data from
different data sources. generate-readme can be used in a GitHub Actions
workflow on a regular schedule to keep the profile up-to-date.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		templatePath := viper.GetString("template")
		generator, err := readme_template.NewGenerator(templatePath)
		if err != nil {
			return err
		}

		outputPath := viper.GetString("output")
		outputFile, err := os.Create(outputPath)
		if err != nil {
			return err
		}

		defer func(outputFile *os.File) {
			_ = outputFile.Close()
		}(outputFile)

		return generator.Generate(outputFile)
	},
}

func init() {
	GenerateCommand.Flags().StringP(
		"template",
		"t",
		"README.template",
		"The path to the template to use to generate the README",
	)
	_ = viper.BindPFlag(
		"template",
		GenerateCommand.Flags().Lookup("template"),
	)

	GenerateCommand.Flags().StringP(
		"output",
		"o",
		"README.md",
		"The path to the README document to generate",
	)
	_ = viper.BindPFlag(
		"output",
		GenerateCommand.Flags().Lookup("output"),
	)
}
