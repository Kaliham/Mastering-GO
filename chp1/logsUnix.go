package main

import (
	"log"
	"log/syslog"
)

func main() {
	sysLogs, err := syslog.New(syslog.LOG_INFO|syslog.LOG_MAIL, "logUnix.go")
	if err != nil {
		log.Fatalln("Error Occured while creating system log output: ", err.Error())
	} else {
		log.SetOutput(sysLogs)
	}
	log.Println("Logging from logUnix.go:", "2024-10-09")
}
