package h

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strconv"
	"sync"

	"github.com/thaibui2308/api-wrapper/config"
	"github.com/thaibui2308/api-wrapper/models"
)

func MapInitials(schoolID string) ([]int, int) {
	var wg sync.WaitGroup
	result := make([]int, 0)
	url := config.BaseURL + schoolID
	numPages := GetNumberOfPages(url)
	intervals := SearchIntervals(numPages)

	result = append(result, 65)

	// If there is only 1 search interval, return an empty array to indicate that there no need for an intials map
	if intervals == 1 {
		return []int{}, numPages
	}

	for i := 1; i <= intervals; i++ {
		wg.Add(1)
		go func(interval int) {
			var pageResponse models.APIResponse
			var pageURL string
			if interval < intervals {
				pageURL = config.PageUrl1 + strconv.Itoa(interval*75) + config.PageUrl2 + schoolID
			} else {
				pageURL = config.PageUrl1 + strconv.Itoa(numPages) + config.PageUrl2 + schoolID
			}

			page, err := http.Get(pageURL)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			responseData, err := ioutil.ReadAll(page.Body)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			json.Unmarshal(responseData, &pageResponse)
			professorList := pageResponse.Professors

			// Get the lastname initial of the last entry
			value := GetAsciiCode(professorList[len(professorList)-1].TLname)
			result = append(result, value)

			wg.Done()
		}(i)
	}
	wg.Wait()
	sort.Ints(result)
	return result, numPages
}
