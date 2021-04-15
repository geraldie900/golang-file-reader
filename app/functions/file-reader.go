package functions

import (
	"io/ioutil"
	"log"
	"regexp"
)

func FileReader(filename string) interface{} {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`\r`)
	content2 := re.ReplaceAllString(string(content), "")

	data := map[string]interface{}{
		"data": string(content2),
	}

	return data
}
