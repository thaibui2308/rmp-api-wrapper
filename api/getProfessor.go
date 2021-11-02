package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"

	h "github.com/thaibui2308/api-wrapper/helpers"
	"github.com/thaibui2308/api-wrapper/models"
)

type SchoolInfo struct {
	ID            string
	NumProfessors int
	NumPages      int
}

func (s *SchoolInfo) FindInstructor(first, last string) (models.Professor, error) {
	// Get the initials ascii array and then page number
	var from, to int
	found := make(chan models.Professor)
	fail := make(chan bool, 2)
	initialMap, numPages := h.MapInitials(s.ID)
	lower, upper := h.SearchInterval(last, initialMap)

	if first == "" && last == "" {
		return models.Professor{}, errors.New("cannot find this professor in the system")

	}

	if lower == 0 && upper == 0 {
		return models.Professor{}, errors.New("cannot find this professor in the system")
	}
	s.NumPages = numPages

	// assign the search bound
	// if there are less than 75 professor at the school
	if len(initialMap) == 0 {
		from = 1
		to = 4
	} else {
		if upper == len(initialMap)-1 {
			from = lower * 75
			to = numPages
		} else if lower == 0 {
			from = 1
			to = 75
		} else {
			from = lower * 75
			to = upper * 75
		}
	}

	// TODO: divide the requests into smaller chunk to cover all of the entries
	go sendRequest(*s, first, last, found, fail, from, to)
	go sendRequest(*s, first, last, found, fail, to+1, to+50)

	for {
		select {
		case m := <-found:
			return m, nil
		case <-fail:
			if len(fail) == 2 {
				return models.Professor{}, errors.New("cannot find this professor in the system")
			}
		}
	}
}

func sendRequest(schoolInfo SchoolInfo, first, last string, found chan models.Professor, fail chan bool, from int, to int) {
	var wg sync.WaitGroup

	// loop through the professor array and run the search func on different thread
	for i := from; i <= to; i++ {
		wg.Add(1)
		go func(id int) {
			var pageResponse models.APIResponse
			page, err := http.Get(pageUrl(id, schoolInfo.ID))
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

			for _, v := range professorList {
				if (v.TLname == last) && (v.TFname == first) {
					found <- v
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	fail <- false
}
