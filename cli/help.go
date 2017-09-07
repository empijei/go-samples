package cli

import (
	"fmt"
	"os"
	"strings"
	"text/template"
	"unicode"
	"unicode/utf8"
)

const docTemplate = `{{.Name | capitalize}}: {{.Short}}

Usage:
	Cli {{.UsageLine}}

{{.Long | trim}}
`

var outw = os.Stderr

var cmdHelp = &Cmd{
	Name:      "help",
	Run:       helpMain,
	UsageLine: "help",
	Short:     "display help information for Cli commands",
	Long:      "",
}

func init() {
	AddCommand(cmdHelp)
}

func helpMain(_ ...string) {
	requestedCmd := "help"
	if len(os.Args) > 1 {
		requestedCmd = os.Args[1]
	}

	if command, err := FindCommand(requestedCmd); err == nil {
		tmpl := template.New("help")
		tmpl.Funcs(template.FuncMap{"trim": strings.TrimSpace, "capitalize": capitalize})
		template.Must(tmpl.Parse(docTemplate))
		_ = tmpl.Execute(outw, command)
	} else {
		fmt.Fprintf(outw, "help: error processing command: %s\n", err.Error())
	}
}

func capitalize(s string) string {
	if s == "" {
		return s
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToTitle(r)) + s[n:]
}
