package model

type Doctor struct {
	Name             string  `json:"name"`
	SpecID           string  `json:"specid"`
	Avatar           string  `json:"avatar"`
	Password         string  `json:"password"`
	Number           string  `json:"number"`
	Online           bool    `json:"online_pay"`
	Experience       int     `json:"experience_years"`
	Address          string  `json:"address"`
	Phone            string  `json:"phone"`
	Weekdays         []bool  `json:"week_days"`
	Comments         int     `json:"comments"`
	Latest_Comment   string  `json:"latest_comment"`
	Latest_Commenter string  `json:"latest_commenter"`
	Average          float64 `json:"average"`
}
