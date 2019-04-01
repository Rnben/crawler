package parser

import (
	"crawler/engine"
	"crawler/model"
	"fmt"
	"regexp"
)

const ageRe = `<div class="m-btn purple" data-v-bff6f798="">(.*?)</div>`

func ParseProfile(contents []byte, name string) engine.ParserResult {
	re := regexp.MustCompile(ageRe)
	profile := model.Profile{}
	profile.Name = name
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		fmt.Println(m)
	}
	result := engine.ParserResult{
		Items: []interface{}{profile},
	}
	return result
}
