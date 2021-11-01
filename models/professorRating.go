package models

type ProfessorRating struct {
	Ratings   []Rating `json:"ratings"`
	Remaining int      `json:"remaining"`
}
