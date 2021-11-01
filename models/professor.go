package models

type Professor struct {
	TDept           string `json:"tDept"`
	TSid            string `json:"tSid"`
	InstitutionName string `json:"institution_name"`
	TFname          string `json:"tFname"`
	TMiddlename     string `json:"tMiddlename"`
	TLname          string `json:"tLname"`
	Tid             int    `json:"tid"`
	TNumRatings     int    `json:"tNumRatings"`
	RatingClass     string `json:"rating_class"`
	ContentType     string `json:"contentType"`
	CategoryType    string `json:"categoryType"`
	OverallRating   string `json:"overall_rating"`
}
