package debug

import (
	"log"
	"os"
	"strings"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Blue = "\033[34m"
var Purple = "\033[35m"

// Log will print info in terminal. Depends of environment and level of log
// will decide of visibility.
// Log levels:
// info - regular info
// warning - warnings :)
// error - errors :(
// debug - this will we logged only when OS ENV variable 'DEBUG' set as 'true'
func Log(level string, s string) {
	if strings.ToUpper(os.Getenv("CI")) == "TRUE" {
		Reset = ""
		Red = ""
		Blue = ""
		Purple = ""
	}
	switch level {
	case "info":
		log.Println("[INFO] " + s)
	case "warning":
		log.Println(Blue + "[WARNING] " + s + Reset)
	case "error":
		log.Println(Red + "[ERROR] " + s + Reset)
	case "debug":
		log.Println(Purple + "[DEBUG] " + Reset + s)
	default:
		log.Println(s)
	}
}
