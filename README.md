# GOALert
This tool will start and monitor a process, waiting it for finish and notifying the exit-code via e-mail message
```
Usage:
  -command string
    	Command to execute
  -version
    	Print version and exit
```

Configuration JSON: $HOME/.config/goalert/conf.json
```json
{
	"username":"username@domain.com",
	"password":"s3cret",
	"from":"username@domain.com",
	"to":"username@domain.com",
	"subject":"unecessary, it is possible to set it in the main file",
	"body":"unecessary, it is possible to set it in the main file",
	"host":"smtp.server_domain.com",
	"port":1337
}
```
