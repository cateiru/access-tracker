package utils_test

import (
	"log"

	"github.com/yuto51942/access-tracker/utils"
)

func IsUrl_test() {
	urls := []string{
		"https//example.com",
		"hoge",
		"",
		"http://nyan.cat",
		"https://youtu.be",
	}

	flag := []bool{
		true,
		false,
		false,
		true,
		true,
	}

	for i, url := range urls {
		if utils.IsUrl(url) != flag[i] {
			log.Fatal(url)
		}
	}
}
