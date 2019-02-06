package loggly

import (
	"goUtils/networking"
	"log"
	"os"
)

const logglyURL string = "http://logs-01.loggly.com/inputs/"

var logglyAPIKey, logglyAPIKeySet = os.LookupEnv("LOGGLY_API_KEY")

// Error log.
func Error(message interface{}) {
	sendOutLogMessage("Error", message)
}

// Warn log.
func Warn(message interface{}) {
	sendOutLogMessage("Warn", message)
}

// Debug log.
func Debug(message interface{}) {
	sendOutLogMessage("Debug", message)
}

// Info log.
func Info(message interface{}) {
	sendOutLogMessage("Info", message)
}

// Trace log.
func Trace(message interface{}) {
	sendOutLogMessage("Trace", message)
}

// Echo Error log.
func ErrorEcho(message interface{}) {
	echoLog("Error", message)
}

// Echo Warn log.
func WarnEcho(message interface{}) {
	echoLog("Warn", message)
}

// Echo Debug log.
func DebugEcho(message interface{}) {
	echoLog("Debug", message)
}

// Echo Info log.
func InfoEcho(message interface{}) {
	echoLog("Info", message)
}

// Echo Trace log.
func TraceEcho(message interface{}) {
	echoLog("Trace", message)
}

// Will print out to the console what the log message is, as well as send it to loggly.
func echoLog(tag string, message interface{}) {
	sendOutLogMessage(tag, message)
	log.Printf("Tag: %s\n %+v\n", tag, message)
}

// Abstracts the message sending to loggly.
func sendOutLogMessage(tag string, message interface{}) {
	if logglyAPIKeySet {
		var url = buildURL(tag)

		switch m := message.(type) {
		case string:
			_, err := networking.Post(url, string(m))
			handleError(err)
		default:
			_, err := networking.PostJson(url, message)
			handleError(err)
		}
	} else {
		log.Fatal(`The "LOGGLY_API_KEY" env variable was not set.  Please set it.`)
	}
}

func handleError(err error) {
	if err != nil {
		log.Println(err)
	}
}

// Builds a URL for the log messages to be sent to.
func buildURL(tag string) string {
	return logglyURL + logglyAPIKey + "/tag/" + tag
}
