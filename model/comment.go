package model

type Comment struct {
	Firstname string `json:"firstname"`
	Doctor_Id string `json:"doctor_id"`
	Reason    string `json:"reason"`
	Desc      string `json:"desc"`
	Stars     int    `json:"stars"`
	Likes     int    `json:"likes"`
	Date      Date   `json:"date"`
}
