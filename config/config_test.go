package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTestConfig() (err error, configFile *os.File) {
	testFileString := `
---

regex_actions:
  normal:
    regex: ".*(норма|norma).*"
  gopstop:
    regex: ".*(gop|гоп).*"

    `
	file, createFileError := ioutil.TempFile("", "config")
	if createFileError != nil {
		return fmt.Errorf("Cannot create temp file: \t%v", createFileError), nil
	}
	if chmodErr := file.Chmod(0644); chmodErr != nil {
		return fmt.Errorf("Cannot chmod temp file: \t%v", chmodErr), nil
	}
	if _, writeStringErr := file.WriteString(testFileString); writeStringErr != nil {
		return fmt.Errorf("Cannot write text to temp file: \t%v", writeStringErr), nil
	}
	return nil, file
}

// TestGetConfig tests config loading
func TestGetConfig(t *testing.T) {
	// Create and load config
	err, configFile := createTestConfig()
	defer os.Remove(configFile.Name())

	assert.Nil(t, err)

	myToken := "myLittleTestToken"
	assert.Nil(t, os.Setenv("FLIBOOBSTIER_TG_TOKEN", myToken))

	config, err := GetConfig(configFile.Name())
	assert.Nil(t, err)

	assert.Equal(t, myToken, config.TgToken)

	// Test words count
	assert.Equal(t, 2, len(config.RegexActions))

	// Test regex compile on "normal" cathch
	assert.Contains(t, config.RegexActions, "normal")
	assert.Contains(t, config.RegexActions, "gopstop")
	hitStr := "Да это же норма, епт"
	missStr := "Нет, это уже пиздец"
	assert.True(t, config.RegexActions["normal"].Regex.MatchString(hitStr))
	assert.False(t, config.RegexActions["normal"].Regex.MatchString(missStr))
}
