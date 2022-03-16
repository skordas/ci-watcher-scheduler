package logging

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var reset = "\033[0m"
var red = "\033[31m"
var blue = "\033[34m"
var purple = "\033[35m"

func Info(s string, a ...interface{}) {
	s = fmt.Sprintf(s, a...)
	log.Println("[INFO] " + s)
}

func Warning(s string, a ...interface{}) {
	setColors()
	s = fmt.Sprintf(s, a...)
	log.Println(blue + "[WARNING] " + s + reset)
}

func Error(s string, a ...interface{}) {
	setColors()
	s = fmt.Sprintf(s, a...)
	log.Println(red + "[ERROR] " + s + reset)
}

func Debug(s string, a ...interface{}) {
	if strings.ToUpper(os.Getenv("DEBUG")) == "TRUE" {
		setColors()
		s = fmt.Sprintf(s, a...)
		log.Println(purple + "[DEBUG] " + reset + s)
	}
}

func setColors() {
	if strings.ToUpper(os.Getenv("CI")) == "TRUE" {
		reset = ""
		red = ""
		blue = ""
		purple = ""
	}
}
