package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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
	configuration = flag.String("configuration", "", "Configuration file [JSON]")
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
		os.Exit(0)
	}
	if *configuration != "" {
		conf = *configuration
		realpath, _ := filepath.Abs(conf)
		if _, err := os.Stat(realpath); os.IsNotExist(err) {
			fmt.Printf("The provided file (%v) does not exist", realpath)
		} else {
			fmt.Printf("conf provided %v", realpath)
		}
		//for testing purposes just print and quit
		os.Exit(2)
	}
	if *command != "" {
		flags.command = strings.Split(*command, " ")
	}
	return flags
}

func Start(cmd []string) string {
	var ExitCodeOut string
	fmt.Printf("---------- COMMAND LINE ----------\n %v \n----------------------------------\n", cmd)
	command := exec.Command(cmd[0], cmd[1:]...)
	startError := command.Start()
	if startError != nil {
		fmt.Println("errore qui")
		log.Fatal(startError)
	}
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
