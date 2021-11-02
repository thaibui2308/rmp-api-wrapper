# REMP
An API wrapper for the two RateMyProfessor public API urls!

You can use this package to find any instructors on the page RateMyProfessor.

### Install 
```
go get -u github.com/thaibui2308/rmp-api-wrapper
```
### Functions
| Name | Signature | Description
| :--- | :--- | :---
| NewInstructor | `NewInstructor(schoolID int, firstname, lastname string) Instructor` | Initiate an Instructor struct
| Find | `(s *Instructor) Find() (models.Professor, error)` | An Instructor struct method to find a specific instructor given their first name and last name
| Ratings | `(s *Instructor) Ratings(num int) ([]models.Rating, error)` | An Instructor struct method to get a specific number of ratings from an instructor


### Structs
```go
type Instructor struct {
	SchoolID    string
	ID          string
	Firstname   string
	Lastname    string
	Institution string
	Found       bool
}
```
```go
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
```
```go
type Rating struct {
	Attendance        string      `json:"attendance"`
	ClarityColor      string      `json:"clarityColor"`
	EasyColor         string      `json:"easyColor"`
	HelpColor         string      `json:"helpColor"`
	HelpCount         int         `json:"helpCount"`
	ID                int         `json:"id"`
	NotHelpCount      int         `json:"notHelpCount"`
	OnlineClass       string      `json:"onlineClass"`
	Quality           string      `json:"quality"`
	RClarity          int         `json:"rClarity"`
	RClass            string      `json:"rClass"`
	RComments         string      `json:"rComments"`
	RDate             string      `json:"rDate"`
	REasy             float64     `json:"rEasy"`
	REasyString       string      `json:"rEasyString"`
	RErrorMsg         interface{} `json:"rErrorMsg"`
	RHelpful          int         `json:"rHelpful"`
	RInterest         string      `json:"rInterest"`
	ROverall          float64     `json:"rOverall"`
	ROverallString    string      `json:"rOverallString"`
	RStatus           int         `json:"rStatus"`
	RTextBookUse      string      `json:"rTextBookUse"`
	RTimestamp        int64       `json:"rTimestamp"`
	RWouldTakeAgain   string      `json:"rWouldTakeAgain"`
	SID               int         `json:"sId"`
	TakenForCredit    string      `json:"takenForCredit"`
	Teacher           interface{} `json:"teacher"`
	TeacherGrade      string      `json:"teacherGrade"`
	TeacherRatingTags []string    `json:"teacherRatingTags"`
	UnUsefulGrouping  string      `json:"unUsefulGrouping"`
	UsefulGrouping    string      `json:"usefulGrouping"`
}
```

### Example
```go
package main

import (
  "github.com/thaibui2308/rmp-api-wrapper"
  "log"
  "fmt"
)

func main() {
  s := remp.NewInstructor(877, "Vadim", "Ponomarenko")
  instructor, err := s.Find()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(instructor)
  
  // Get two ratings from this instructor
  ratings, err := Ratings(2)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(ratings)
}
```
