package models

type ProfileManager struct {
	Id           int    `json:"id"`
	AuthId       int    `json:"manager_id"`
	Name         string `json:"name"`
	userImage    string `json:"user_image"`
	companyName  string `json:"company_name"`
	companyImage string `json:"company_image"`
}

type ProfileManagerAuth struct {
	Id           int    `json:"id"`
	Email        string `json:"email"`
	AuthId       int    `json:"manager_id"`
	Name         string `json:"name"`
	userImage    string `json:"user_image"`
	companyName  string `json:"company_name"`
	companyImage string `json:"company_image"`
}

type ProfileManagerClaims struct {
	Id     int    `json:"id"`
	AuthId int    `json:"manager_id"`
	Email  string `json:"email"`
}
