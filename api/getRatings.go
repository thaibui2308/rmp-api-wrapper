package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/thaibui2308/api-wrapper/models"
)

func GetRatings(tid string) []models.Rating {
	var ratings models.ProfessorRating

	response, err := http.Get("https://www.ratemyprofessors.com/paginate/professors/ratings?tid=" + tid + "&filter=&courseCode=&page=1")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(responseData, &ratings)
	ratingList := ratings.Ratings
	return ratingList
}
