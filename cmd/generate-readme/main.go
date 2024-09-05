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

// This program implements a GitHub Action that will generate a GitHub user
// or organization profile README.md document using a template. This program
// uses the Go text templating engine to generate the GitHub profile README.md
// document. Profile authors can inject data from different data files or
// remote sources to complement their GitHub profiles.

package main

import (
	"github.com/spf13/cobra"
	"log"
)

var rootCommand = &cobra.Command{
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
}

func main() {
	if err := rootCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
