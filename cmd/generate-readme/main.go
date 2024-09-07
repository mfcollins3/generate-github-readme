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
	"github.com/spf13/viper"
	"log"
)

func init() {
	viper.SetDefault("template", "README.template")
	viper.SetDefault("output", "README.md")

	viper.AutomaticEnv()
	viper.SetEnvPrefix("INPUT")

	rootCommand.Flags().StringP(
		"template",
		"t",
		"README.template",
		"The path to the template to use to generate the README",
	)
	_ = viper.BindPFlag("template", rootCommand.Flags().Lookup("template"))

	rootCommand.Flags().StringP(
		"output",
		"o",
		"README.md",
		"The path to the README document to generate",
	)
	_ = viper.BindPFlag("output", rootCommand.Flags().Lookup("output"))
}

func main() {
	if err := rootCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
