package getopt

import (
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func mapifyConfig(environment []string) (envArray map[string]string) {
	envArray = make(map[string]string)

	for _, cur := range environment {
		envVar := strings.Split(cur, "=")
		if len(envVar) > 1 {
			envArray[strings.TrimSpace(envVar[0])] = strings.TrimSpace(envVar[1])
		}
	}

	return
}

func readConfigFile(path string) (configEntries []string, err *GetOptError) {
	validConfigEntry := regexp.MustCompile("^[A-z0-9_.,]+=.*$")

	content, ioErr := ioutil.ReadFile(path)
	contentStringified := string(content)

	if ioErr != nil {
		err = &GetOptError{ConfigFileNotFound, ioErr.Error()}
	} else {
		for _, line := range strings.Split(contentStringified, "\n") {
			if validConfigEntry.MatchString(line) {
				configEntries = append(configEntries, line)
			}
		}
	}

	return
}

func processConfigFile(path string) (err *GetOptError) {
	configEntries, err := readConfigFile(path)
	if err == nil {
		for key, value := range mapifyConfig(configEntries) {
			if os.Getenv(key) == "" {
				os.Setenv(key, value)
			}
		}
	}

	return
}
