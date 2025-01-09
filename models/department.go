package models

type Department struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	ProfileId	string `json:"profileId"`
}