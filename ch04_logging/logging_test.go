package logging

import (
	"io"
	"log"
	"log/syslog"
	"os"
	"strconv"
	"testing"
)

func TestBasic(t *testing.T) {
	str := "abcdefghi"
	_, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Println("Cannot parse string:", err)
	}
}

func TestFlag(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Some event happened")
}

func TestPrefix(t *testing.T) {
	log.SetPrefix("INFO ")
	log.Println("Some event happened")
	log.Println("Another event happened")
}

func TestOutputFile(t *testing.T) {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(file)
	log.Println("Some event happened")
	log.Println("Another event happened")
}

func TestOutputMultiWriter(t *testing.T) {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := io.MultiWriter(os.Stderr, file)
	log.SetOutput(writer)
	log.Println("Some event happened")
	log.Println("Another event happened")

}

func TestLogLevel(t *testing.T) {
	info.Println("Some informational event happened")
	debug.Println("Some debugging event happened")
}

func TestOutputFunc(t *testing.T) {
	err := log.Output(2, "this is a test log")
	if err != nil {
		t.Error("Cannot write to log", err)
	}
}

func TestSyslog(t *testing.T) {

	logger, err := syslog.NewLogger(syslog.LOG_NOTICE, log.LstdFlags)
	if err != nil {
		t.Error("error:", err)
	}
	logger.Print("Hello Logs!")
}
