package main

func main() {
	var cmd Commands
	cmd = Init()
	ExitCode := Start(cmd)
	//SendMail(ExitCode)
	_ = MockSendMail(MockMail{ExitCode, "from", "to", "subject"})
}
