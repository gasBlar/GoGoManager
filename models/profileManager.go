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
	Id           int            `json:"id"`
	Email        sql.NullString `json:"email"`
	AuthId       int            `json:"managerId"`
	Name         sql.NullString `json:"name"`
	UserImage    sql.NullString `json:"UserImage"`
	CompanyName  sql.NullString `json:"CompanyName"`
	CompanyImage sql.NullString `json:"CompanyImage"`
}

type ProfileManagerClaims struct {
	Id     int    `json:"id"`
	AuthId int    `json:"managerId"`
	Email  string `json:"email"`
}

type ProfileManagerUpdateRequest struct {
	Email           string `json:"email" validate:"required,email"`
	Name            string `json:"name" validate:"omitempty,min=4,max=52"`
	UserImageUri    string `json:"UserImageUri" validate:"omitempty"`
	CompanyName     string `json:"CompanyName" validate:"omitempty,min=4,max=52"`
	CompanyImageUri string `json:"CompanyImageUri" validate:"omitempty"`
}

type ProfileManagerResponse struct {
	Email        string `json:"email"`
	Name         string `json:"name"`
	UserImage    string `json:"UserImageUri"`
	CompanyName  string `json:"CompanyName"`
	CompanyImage string `json:"CompanyImageUri"`
}

func (p *ProfileManagerAuth) ToResponse() ProfileManagerResponse {
	return ProfileManagerResponse{
		Email:        NullStringToString(p.Email),
		Name:         NullStringToString(p.Name),
		UserImage:    NullStringToString(p.UserImage),
		CompanyName:  NullStringToString(p.CompanyName),
		CompanyImage: NullStringToString(p.CompanyImage),
	}
}

func NullStringToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}
