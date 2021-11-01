package h

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/thaibui2308/api-wrapper/models"
)

// Get the number of response page
func GetNumberOfPages(url string) int {
	var pageResponse models.APIResponse

	response, err := http.Get(url)
	if err != nil {
		return -1
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return -1
	}
	json.Unmarshal(data, &pageResponse)
	result := (pageResponse.SearchResultsTotal / 20) + 1
	return result
}

// Each search interval contains 100 pages
func SearchIntervals(page int) int {
	result := page/75 + 1
	return result
}

// Get the ascii code of the last name first letter
func GetAsciiCode(letter string) int {
	if letter == "" {
		return 0
	}
	firstLetter := strings.ToUpper(letter[0:1])
	runes := []rune(firstLetter)
	asciiVal := int(runes[0])
	return asciiVal
}
