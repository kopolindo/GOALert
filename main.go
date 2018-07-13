package main

import (
	"github.com/kopolindo/GOAlert/mailert"
)

func main() {
	var flags Flag
	headers := make(map[string]string)
	flags = Init()
	ExitCode := Start(flags.command)
	mailert.SetHeaders(true, headers)
	mailert.SendMail(ExitCode)
}
