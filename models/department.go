package models

type Department struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	ProfileId string `json:"profileId"`
}

type DepartmentPatch struct {
	Id        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	ProfileId string `json:"profileId,omitempty"`
}
