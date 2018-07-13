package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/alecthomas/template"
)

type Version struct {
	Ver   string
	Build string
}

var (
	ActualVersion = Version{"0.1", "dev"}
	version       = flag.Bool("version", false, "Print version and exit")
	command       = flag.String("command", "", "Command to execute")
)

type Flag struct {
	command []string
}

func PrintBanner() {
	dat, err := ioutil.ReadFile("banner")
	if err != nil {
		fmt.Println(err)
	}
	banner := string(dat)
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

func Init() Flag {
	flag.Parse()
	var flags Flag
	if flag.NFlag() == 0 {
		Usage()
		os.Exit(0)
	}
	if *version {
		PrintBanner()
	}
	if *command != "" {
		flags.command = strings.Split(*command, " ")
	}
	return flags
}

func Start(cmd []string) string {
	command := exec.Command(cmd[0], cmd[1:]...)
	startError := command.Start()
	if startError != nil {
		log.Fatal(startError)
	}
	log.Println("Waiting for command to finish...")
	ExitCode := command.Wait()
	log.Printf("Command finished with exit-code: %v", ExitCode)
	return ExitCode.Error()
}
