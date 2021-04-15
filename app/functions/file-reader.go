/*
=====================================
;	Author : Geraldie Tanu Saputra
;	Email  : geraldie.saputra@soluix.ai
;	Date   : 15-04-2021
=====================================
*/

package functions

import (
	"io/ioutil"
	"log"
	"regexp"

	Utils "golang-file-reader/utils"
)

func FileReader(filename string) interface{} {
	path := Utils.FOLDER_PATH

	content, err := ioutil.ReadFile(path + filename)

	if err != nil {
		log.Println("::[ERROR] - ", err)
	}

	re := regexp.MustCompile(`\r`)
	contentConverted := re.ReplaceAllString(string(content), "")

	data := map[string]interface{}{
		"content": string(contentConverted),
	}

	return data
}
