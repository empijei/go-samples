package cli

import (
	"flag"
	"log"
	"os"
	"text/template"

	"github.com/empijei/go-samples/cli/l"
)

func init() {

	// Setup fallback version and commit in case Cli wasn't "properly" compiled
	if len(Version) == 0 {
		Version = "Unknown, please compile Cli with 'make'"
	}
	if len(Commit) == 0 {
		Commit = "Unknown, please compile Cli with 'make'"
	}
	if len(Build) == 0 {
		Build = "Debug"
	}

}

// ASCII art generated with figlet
const banner = `      _ _ 
  ___| (_)  
 / __| | |  Version: {{.Version}}
| (__| | |  Commit:  {{.Commit}}
 \___|_|_|  Build:   {{.Build}}

`

var (
	//Version is taken by the build flags, represent current version as
	//<major>.<minor>.<patch>
	Version string

	//Commit is the output of `git rev-parse --short HEAD` at the moment of the build
	Commit string

	//Either "Release", "Testing" or "Debug"
	Build string
)

func Printbanner() {
	tmpl := template.New("banner")
	template.Must(tmpl.Parse(banner))
	_ = tmpl.Execute(os.Stderr, struct{ Version, Commit, Build string }{Version, Commit, Build})
}

func Init() {
	if Build == "Release" {
		l.CurLevel = l.Level_Info
		l.SetFlags(log.Ltime)
	}

	stdoutinfo, err := os.Stdout.Stat()
	if err == nil && stdoutinfo.Mode()&os.ModeCharDevice == 0 {
		// Output is a pipe, turn off colors
		l.Color = false
	} else {
		// Output is to terminal, print banner
		Printbanner()
	}

	if len(CliCommands) > 2 {
		initCommands()
	} else {
		flag.Parse()
	}
}
