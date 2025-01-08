package models

import (
	"database/sql"
)

type ProfileManager struct {
	Id           int    `json:"id"`
	AuthId       int    `json:"manager_id"`
	Name         string `json:"name"`
	UserImage    string `json:"user_image"`
	CompanyName  string `json:"company_name"`
	CompanyImage string `json:"company_image"`
}

type ProfileManagerAuth struct {
	Id           int            `json:"id"`
	Email        sql.NullString `json:"email"`
	AuthId       int            `json:"manager_id"`
	Name         sql.NullString `json:"name"`
	UserImage    sql.NullString `json:"user_image"`
	CompanyName  sql.NullString `json:"company_name"`
	CompanyImage sql.NullString `json:"company_image"`
}

type ProfileManagerClaims struct {
	Id     int    `json:"id"`
	AuthId int    `json:"manager_id"`
	Email  string `json:"email"`
}

type ProfileManagerResponse struct {
	Email        string `json:"email"`
	Name         string `json:"name"`
	UserImage    string `json:"user_image"`
	CompanyName  string `json:"company_name"`
	CompanyImage string `json:"company_image"`
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
