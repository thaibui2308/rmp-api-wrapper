package rmp

import (
	"errors"
	"strconv"

	"github.com/thaibui2308/api-wrapper/api"
	"github.com/thaibui2308/api-wrapper/models"
)

type Instructor struct {
	SchoolID    string
	ID          string
	Firstname   string
	Lastname    string
	Institution string
	Found       bool
}

func NewInstructor(schoolID int, firstname, lastname string) Instructor {
	sID := strconv.Itoa(schoolID)
	return Instructor{
		SchoolID:  sID,
		Firstname: firstname,
		Lastname:  lastname,
	}
}

func (s *Instructor) Ratings(num int) ([]models.Rating, error) {
	if s.ID == "" {
		return []models.Rating{}, errors.New("teacher id is not provided")
	}

	ratings := api.GetRatings(s.ID)
	if num <= 0 || num > len(ratings) {
		return []models.Rating{}, errors.New("invalid number of ratings")
	}

	return ratings[:num], nil
}

func (s *Instructor) Find() (models.Professor, error) {
	if s.SchoolID == "" {
		return models.Professor{}, errors.New("school id is not provided")
	}
	schoolInfo := api.SchoolInfo{
		ID: s.SchoolID,
	}
	professor, err := schoolInfo.FindInstructor(s.Firstname, s.Lastname)
	if err != nil {
		s.Found = false
		return professor, err
	}

	s.Found = true
	s.ID = strconv.Itoa(professor.Tid)
	s.Institution = professor.InstitutionName

	return professor, err
}
