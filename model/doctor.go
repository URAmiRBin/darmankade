package model

type Doctor struct {
	Name       string `json:"name"`
	SpecID     string `json:"spec"`
	Avatar     string `json:"avatar"`
	Password   string `json:"avatar"`
	Number     string `json:"number"`
	Online     bool   `json:"online_pay"`
	Experience int    `json:"experience_years"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
	Weekdays   []bool `json:"week_days"`
}
