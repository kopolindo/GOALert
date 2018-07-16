package main

func main() {
	var flags Flag
	flags = Init()
	ExitCode := Start(flags.command)
	SendMail(ExitCode)
}
