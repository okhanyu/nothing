package common

type Page struct {
	PageNumber int `json:"page" form:"page"`
	PageSize   int `json:"size" form:"size"`
}
