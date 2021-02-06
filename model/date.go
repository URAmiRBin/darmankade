package model

var MonthMapper = map[int]string{
	1:  "فروردین",
	2:  "اردیبهشت",
	3:  "خرداد",
	4:  "تیر",
	5:  "مرداد",
	6:  "شهریور",
	7:  "مهر",
	8:  "آبان",
	9:  "آذر",
	10: "دی",
	11: "بهمن",
	12: "اسفند",
}

type Date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

func (d Date) String() string {
	return string(d.Day) + MonthMapper[d.Month] + string(d.Year)
}
