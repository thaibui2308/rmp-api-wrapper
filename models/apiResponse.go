package models

type APIResponse struct {
	Professors         []Professor `json:"professors"`
	SearchResultsTotal int         `json:"searchResultsTotal"`
	Remaining          int         `json:"remaining"`
	Type               string      `json:"type"`
}
