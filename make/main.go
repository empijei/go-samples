package main

import (
	"html/template"
	"os"
)

const banner = `

  Version: {{.Version}}
  Commit:  {{.Commit}}
  Build:   {{.Build}}

 
`

func printbanner() {
	tmpl := template.New("banner")
	template.Must(tmpl.Parse(banner))
	_ = tmpl.Execute(os.Stderr, struct{ Version, Commit, Build string }{Version, Commit, Build})
}

func init() {
	// Setup fallback version and commit in case wapty wasn't "properly" compiled
	if len(Version) == 0 {
		Version = "Unknown, please compile wapty with 'make'"
	}
	if len(Commit) == 0 {
		Commit = "Unknown, please compile wapty with 'make'"
	}
	if len(Build) == 0 {
		Build = "Debug"
	}
}

var (
	//Version is taken by the build flags, represent current version as
	//<major>.<minor>.<patch>
	Version string

	//Commit is the output of `git rev-parse --short HEAD` at the moment of the build
	Commit string

	//Either "Release", "Testing" or "Debug"
	Build string
)

func main() {
	printbanner()
}
