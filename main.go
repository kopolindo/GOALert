package main

func main() {
	var cmd Commands
	cmd = Init()
	ExitCode := Start(cmd)
	SendMail(ExitCode)
}
