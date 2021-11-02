package api

import (
	"strconv"

	"github.com/thaibui2308/api-wrapper/config"
)

func pageUrl(pageNumber int, schoolID string) string {
	url := config.PageUrl1 + strconv.Itoa(pageNumber) + config.PageUrl2 + schoolID
	return url
}
