package template

import (
	"github.com/fatih/color"
	"strings"
)

var bold = color.New(color.Bold)

var colorizedHelpTemplate = []string{
	"",
	"    " + bold.Sprintf("{{.HelpName}} ") + "{{if .VisibleFlags}}[global options]{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}{{if .Version}}{{if not .HideVersion}}",
	"",
	"    " + bold.Sprintf("Version:"),
	"        {{.Version}}{{end}}{{end}}{{if .Description}}",
	"",
	"    " + bold.Sprintf("Description:"),
	"        {{.Description}}{{end}}{{if len .Authors}}",
	"",
	"    " + bold.Sprintf("Author{{with $length := len .Authors}}{{if ne 1 $length}}S{{end}}{{end}}:"),
	"        {{range $index, $author := .Authors}}{{if $index}}",
	"        {{end}}{{$author}}{{end}}{{end}}{{if .VisibleCommands}}",
	"",
	"    " + bold.Sprintf("Commands:") + "{{range .VisibleCategories}}{{if .Name}}",
	"        {{.Name}}:{{end}}{{range .VisibleCommands}}",
	"        {{join .Names \", \"}}{{\"\t\"}}{{.Usage}}{{end}}{{end}}{{end}}{{if .VisibleFlags}}",
	"",
	"    " + bold.Sprintf("Global Options:"),
	"        {{range $index, $option := .VisibleFlags}}{{if $index}}",
	"        {{end}}{{$option}}{{end}}{{end}}{{if .Copyright}}",
	"",
	"    " + bold.Sprintf("Copyright:"),
	"   {{.Copyright}}{{end}}",
	"\n",
}

var colorizedCommandHelpTemplate = []string{
	"",
	"    " + bold.Sprintf("{{.HelpName}}") + " - {{.Usage}}",
	"",
	"    " + bold.Sprintf("Usage:"),
	"        {{.HelpName}}{{if .VisibleFlags}} [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{if .Category}}",
	"",
	"    " + bold.Sprintf("Category:"),
	"        {{.Category}}{{end}}{{if .Description}}",
	"",
	"    " + bold.Sprintf("Description:"),
	"        {{.Description}}{{end}}{{if .VisibleFlags}}",
	"",
	"    " + bold.Sprintf("Options:"),
	"        {{range .VisibleFlags}}{{.}}",
	"        {{end}}{{end}}",
	"\n",
}

var ColorizedCommandHelpTemplate = strings.Join(colorizedCommandHelpTemplate, "\n")
var ColorizedHelpTemplate = strings.Join(colorizedHelpTemplate, "\n")
