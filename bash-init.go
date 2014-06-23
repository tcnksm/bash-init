package main

import (
	"fmt"
	flag "github.com/dotcloud/docker/pkg/mflag"
	"log"
	"os"
	"strings"
	"text/template"
)

var mainTemplate = template.Must(parseAsset("main", "templates/main.tmpl"))
var readmeTemplate = template.Must(parseAsset("readme", "templates/README.tmpl"))

var readmeMd = Source{
	Name:     "README.md",
	Template: *readmeTemplate,
}

type application struct {
	Name, Author, Email string
	HasSubCommand       bool
	SubCommands         []subCommand
}

type subCommand struct {
	Name string
}

func parseAsset(name string, path string) (*template.Template, error) {
	src, err := Asset(path)
	if err != nil {
		return nil, err
	}

	return template.New(name).Parse(string(src))
}

func defineApplication(appName string, inputSubCommands []string) application {

	hasSubCommand := false
	if inputSubCommands[0] != "" {
		hasSubCommand = true
	}

	return application{
		Name:          appName,
		Author:        GitConfig("user.name"),
		Email:         GitConfig("user.email"),
		HasSubCommand: hasSubCommand,
		SubCommands:   defineSubCommands(inputSubCommands),
	}
}

func defineSubCommands(inputSubCommands []string) []subCommand {

	var subCommands []subCommand

	if inputSubCommands[0] == "" {
		return subCommands
	}

	for _, name := range inputSubCommands {
		subCommand := subCommand{
			Name: name,
		}
		subCommands = append(subCommands, subCommand)
	}

	return subCommands
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func showVersion() {
	fmt.Fprintf(os.Stderr, "cli-init v%s\n", Version)
}

func showHelp() {
	fmt.Fprintf(os.Stderr, helpText)
}

func main() {

	var (
		flVersion     = flag.Bool([]string{"v", "-version"}, false, "Print version information and quit")
		flHelp        = flag.Bool([]string{"h", "-help"}, false, "Print this message and quit")
		flDebug       = flag.Bool([]string{"-debug"}, false, "Run as DEBUG mode")
		flSubCommands = flag.String([]string{"s", "-subcommands"}, "", "Conma-seplated list of sub-commands to build")
		flReadme      = flag.Bool([]string{"r", "-readme"}, false, "Include README.md file")
		flForce       = flag.Bool([]string{"f", "-force"}, false, "Overwrite application without prompting")
	)

	flag.Parse()

	if *flHelp {
		showHelp()
		os.Exit(0)
	}

	if *flVersion {
		showVersion()
		os.Exit(0)
	}

	if *flDebug {
		os.Setenv("DEBUG", "1")
		debug("Run as DEBUG mode")
	}

	inputSubCommands := strings.Split(*flSubCommands, ",")
	debug("inputSubCommands:", inputSubCommands)

	appName := flag.Arg(0)
	debug("appName:", appName)

	if appName == "" {
		fmt.Fprintf(os.Stderr, "Application name must not be blank\n")
		os.Exit(1)
	}

	// Define Application
	application := defineApplication(appName, inputSubCommands)

	// Create <appName>.sh
	mainSh := Source{
		Name:     appName + ".sh",
		Template: *mainTemplate,
	}

	removed, err := mainSh.safeRemove(*flForce)
	assert(err)

	if removed {
		err = mainSh.generate(application)
		assert(err)
	}

	if *flReadme {
		removed, err := readmeMd.safeRemove(*flForce)
		assert(err)

		if removed {
			err = readmeMd.generate(application)
			assert(err)
		}
	}

	os.Exit(0)
}

const helpText = `Usage: bash-init [options] [application]

bash-init is the easy way to start building command-line app.

Options:

  -s="", --subcommands=""    Comma-separated list of sub-commands to build
  -f, --force                Overwrite application without prompting
	-r, --readme               Include README.md
  -h, --help                 Print this message and quit
  -v, --version              Print version information and quit
  --debug=false              Run as DEBUG mode

Example:

  $ bash-init todo
  $ bash-init -s add,list,delete todo
`
