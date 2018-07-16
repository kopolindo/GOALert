package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	gomail "gopkg.in/gomail.v2"
)

type Conf struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	From     string  `json:"from"`
	To       string  `json:"to"`
	Subject  string  `json:"subject"`
	Body     string  `json:"body"`
	Host     string  `json:"host"`
	Port     float64 `json:"port"`
}

var (
	host       string
	port       int
	username   string
	password   string
	parameters Conf
	headers    = map[string]string{
		"From":    "",
		"To":      "",
		"Subject": "",
	}
	HOME          = os.Getenv("HOME")
	configuration = filepath.Join(HOME, ".config/goalert/conf.json")
)

func GetConf(fname string) Conf {
	var configuration Conf
	file, err := os.Open(fname)
	if err != nil {
		fmt.Printf("Error opening configuration file %v (%v)\n", fname, err)
		os.Exit(1)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	decodingError := decoder.Decode(&configuration)
	if decodingError != nil {
		fmt.Printf("Decoding error: %s\n", decodingError)
	}
	file.Close()
	return configuration
}

func setParametersFromConf() {
	parameters := GetConf(configuration)
	headers["From"] = parameters.From
	headers["To"] = parameters.To
	headers["Subject"] = "GOALert_TEST"
	host = parameters.Host
	port = int(parameters.Port)
	username = parameters.Username
	password = parameters.Password
}

func SendMail(body string) {
	setParametersFromConf()
	message := gomail.NewMessage()
	message.SetHeader("From", headers["From"])
	message.SetHeader("To", headers["To"])
	message.SetHeader("Subject", headers["Subject"])
	message.SetBody("text/html", body)
	//fmt.Printf("From: %s\nTo: %s\nSubject: %s\nBody: %s\n", headers["From"], headers["To"], headers["Subject"], body)
	dialer := gomail.NewDialer(host, port, username, password)
	if DialErr := dialer.DialAndSend(message); DialErr != nil {
		fmt.Println(DialErr)
	}
}

func SetHeaders(heads map[string]string) {
	headers["From"] = heads["From"]
	headers["To"] = heads["To"]
	headers["Subject"] = heads["Subject"]
}
