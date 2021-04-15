/*
=====================================
;	Author : Geraldie Tanu Saputra
;	Email  : geraldie.saputra@soluix.ai
;	Date   : 15-04-2021
=====================================
*/

package functions

import (
	"bufio"
	"io/ioutil"
	"log"
	"regexp"

	"github.com/gofiber/fiber/v2"
)

func FileReceiver(c *fiber.Ctx) interface{} {
	file, err := c.FormFile("file")

	if err != nil {
		log.Println("::[ERROR] - ", err)
	}

	fileheader, _ := file.Open()
	reader := bufio.NewReader(fileheader)
	content, _ := ioutil.ReadAll(reader)

	re := regexp.MustCompile(`\r`)
	contentConverted := re.ReplaceAllString(string(content), "")

	data := map[string]interface{}{
		"content": string(contentConverted),
	}

	return data

}
