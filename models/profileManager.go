package models

import (
	"database/sql"
)

type ProfileManager struct {
	Id           int    `json:"id"`
	AuthId       int    `json:"managerId"`
	Name         string `json:"name"`
	UserImage    string `json:"UserImage"`
	CompanyName  string `json:"CompanyName"`
	CompanyImage string `json:"CompanyImage"`
}

type ProfileManagerAuth struct {
	Id              int            `json:"id"`
	Email           sql.NullString `json:"email"`
	AuthId          int            `json:"managerId"`
	Name            sql.NullString `json:"name"`
	UserImageUri    sql.NullString `json:"UserImage"`
	CompanyName     sql.NullString `json:"CompanyName"`
	CompanyImageUri sql.NullString `json:"CompanyImage"`
}

type ProfileManagerClaims struct {
	Id     int    `json:"id"`
	AuthId int    `json:"managerId"`
	Email  string `json:"email"`
}

type ProfileManagerUpdateRequest struct {
	Email           string `json:"email" validate:"required,email"`
	Name            string `json:"name" validate:"min=4,max=52,nonempty"`
	UserImageUri    string `json:"userImageUri" validate:"nonempty,uri"`
	CompanyName     string `json:"companyName" validate:"nonempty,min=8,max=32"`
	CompanyImageUri string `json:"companyImageUri" validate:"nonempty,uri"`
}

type ProfileManagerResponse struct {
	Email           string `json:"email"`
	Name            string `json:"name"`
	UserImageUri    string `json:"userImageUri"`
	CompanyName     string `json:"companyName"`
	CompanyImageUri string `json:"companyImageUri"`
}

func (p *ProfileManagerAuth) ToResponse() ProfileManagerResponse {
	return ProfileManagerResponse{
		Email:           NullStringToString(p.Email),
		Name:            NullStringToString(p.Name),
		UserImageUri:    NullStringToString(p.UserImageUri),
		CompanyName:     NullStringToString(p.CompanyName),
		CompanyImageUri: NullStringToString(p.CompanyImageUri),
	}
}

func NullStringToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}
