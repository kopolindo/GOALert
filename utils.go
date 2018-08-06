package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

const banner = `
   __________  ___    __              __ 
  / ____/ __ \/   |  / /   ___  _____/ /_
 / / __/ / / / /| | / /   / _ \/ ___/ __/
/ /_/ / /_/ / ___ |/ /___/  __/ /  / /_  
\____/\____/_/  |_/_____/\___/_/   \__/  

|Version:	{{.Ver}}
|Commit:	{{.Commit}}
|Build:		{{.Build}}
`

var (
	version       = flag.Bool("version", false, "Print version and exit")
	command       = flag.String("command", "", "Command to execute")
	conf          = flag.String("conf", "", "Configuration file [JSON]")
	Version       string
	Commit        string
	Build         string
	ActualVersion = struct{ Ver, Commit, Build string }{Version, Commit, Build}
)

type Commands []string

func PrintBanner() {
	templ := template.New("banner")
	template.Must(templ.Parse(banner))
	_ = templ.Execute(os.Stderr, ActualVersion)
}

func Usage() {
	PrintBanner()
	fmt.Println("\nUsage flags:")
	flag.PrintDefaults()
	return
}

func Init() Commands {
	flag.Usage = func() {
		PrintBanner()
		flag.PrintDefaults()
		os.Exit(0)
	}
	flag.Parse()
	var cmd Commands
	if flag.NFlag() == 0 {
		Usage()
		os.Exit(0)
	}
	if *version {
		flag.Usage()
	}
	if *conf != "" {
		configuration = *conf
		realpath, _ := filepath.Abs(configuration)
		if _, err := os.Stat(realpath); os.IsNotExist(err) {
			fmt.Printf("The provided file (%v) does not exist", realpath)
		} else {
			fmt.Printf("configuration provided %v", realpath)
		}
		//for testing purposes just print and quit
		os.Exit(2)
	}
	if *command != "" {
		cmd = strings.Split(*command, " ")
	}
	return cmd
}

func Start(cmd []string) string {
	var ExitCodeOut string
	fmt.Printf("---------- COMMAND LINE ----------\n %v \n----------------------------------\n", cmd)
	command := exec.Command(cmd[0], cmd[1:]...)
	//create stdout pipeline
	stdout, _ := command.StdoutPipe()
	//start execution and error handling
	startError := command.Start()
	if startError != nil {
		log.Fatal(startError)
	}
	//scanning stdout pipeline
	scanOut := bufio.NewScanner(stdout)
	//printint stdout pipeline buffer
	for scanOut.Scan() {
		out := scanOut.Text()
		fmt.Println(out)
	}
	//waiting for exiting and exit-code
	log.Println("Waiting for command to finish...")
	ExitCode := command.Wait()
	if ExitCode == nil {
		ExitCodeOut = "0"
	} else {
		ExitCodeOut = ExitCode.Error()
	}
	log.Printf("Command finished with exit-code: %v", ExitCodeOut)
	return ExitCodeOut
}
