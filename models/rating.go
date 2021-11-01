package models

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
