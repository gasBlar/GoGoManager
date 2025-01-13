package models

type Department struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	ProfileId int    `json:"profileId"`
}

type DepartmentPatch struct {
	Id        int     `json:"id,omitempty"`
	Name      string  `json:"name,omitempty"`
	ProfileId *int    `json:"profileId,omitempty"` // Gunakan pointer untuk membedakan null
}
