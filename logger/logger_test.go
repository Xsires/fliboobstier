package logger

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

// TestInit tests logger init
func TestInit(t *testing.T) {
	file, createFileError := ioutil.TempFile("", "config")
	if createFileError != nil {
		t.Fatalf("Cannot create temp file: \t%v", createFileError)
	}
	logPath := file.Name()
	defer os.Remove(file.Name())
	Init(logPath)

	logText := "This_is_my_error"
	Logger.Error(logText)

	logFileText, err := ioutil.ReadFile(logPath)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(logFileText), logText) {
		t.Fatal("Cannot find log record on log file")
	}
}
