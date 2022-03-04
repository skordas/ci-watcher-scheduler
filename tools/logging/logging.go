package logging

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Blue = "\033[34m"
var Purple = "\033[35m"

func Info(s string, a ...interface{}) {
	s = fmt.Sprintf(s, a...)
	log.Println("[INFO] " + s)
}

func Warning(s string, a ...interface{}) {
	SetColors()
	s = fmt.Sprintf(s, a...)
	log.Println(Blue + "[WARNING] " + s + Reset)
}

func Error(s string, a ...interface{}) {
	SetColors()
	s = fmt.Sprintf(s, a...)
	log.Println(Red + "[ERROR] " + s + Reset)
}

func Debug(s string, a ...interface{}) {
	if strings.ToUpper(os.Getenv("DEBUG")) == "TRUE" {
		SetColors()
		s = fmt.Sprintf(s, a...)
		log.Println(Purple + "[DEBUG] " + Reset + s)
	}
}

func SetColors() {
	if strings.ToUpper(os.Getenv("CI")) == "TRUE" {
		Reset = ""
		Red = ""
		Blue = ""
		Purple = ""
	}
}
